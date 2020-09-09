package initialization

import (
	"github.com/novliang/lurkers/context"
	"github.com/novliang/lurkers/response"
	"github.com/novliang/lurkers/validator"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func Echo() *echo.Echo {
	e := echo.New()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &validator.Validator{
		Validator: validator2.New(),
	}

	e.HTTPErrorHandler = HttpErrorHandler

	//Extending
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			think := &context.Context{Context: c, JsonParams: nil, RequestTime: time.Now().Unix()}
			return handlerFunc(think)
		}
	})

	return e
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  string
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if h, o := he.Message.(string); o {
			msg = h
		} else {
			msg = "服务器出错!"
		}
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			r := new(response.Response)
			r.Message = msg
			r.Code = code
			err := c.JSON(200, r)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
