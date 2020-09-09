package context

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"github.com/novliang/lurkers/response"
	"net/http"
)

type Context struct {
	echo.Context
	JsonParams  map[string]interface{}
	RequestTime int64
}

func (c *Context) JsonParam(i string) interface{} {
	if c.JsonParams == nil {
		param := map[string]interface{}{}
		body := c.Request().Body
		result, err := ioutil.ReadAll(body)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(result))

		if err != nil {
			c.JsonParams = param
			return ""
		}

		err = json.Unmarshal(result, &param)
		if err != nil {
			c.JsonParams = param
			return ""
		}
		c.JsonParams = param
	}

	v, ok := c.JsonParams[i]
	if ok {
		return v
	}

	return nil
}

func (c *Context) Out(i interface{}) error {

	//New Encoder
	e := json.NewEncoder(c.Context.Response())

	//Get Header
	header := c.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	}

	//Set Header Code
	c.Response().WriteHeader(http.StatusOK)

	//Appointment Response
	ar := new(response.Response)
	ar.Code = 0
	ar.Message = "success"
	ar.Data = i

	return e.Encode(ar)
}

//
//func (c *Context) GetUserID() (uint, error) {
//	// 设置分组ID
//	token, err := global.GIANT_OAUTH_SERVER.ValidationBearerToken(c.Request())
//	if err != nil {
//		return 0, err
//	}
//	UserID := token.GetUserID()
//	if UserID == "" {
//		return 0, errors.New("获取用户信息失败")
//	}
//
//	uid, err := strconv.Atoi(UserID)
//	if err != nil {
//		return 0, err
//	}
//
//	return uint(uid), nil
//
//}
