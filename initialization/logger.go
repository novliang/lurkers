package initialization

import (
	"fmt"
	"github.com/novliang/lurkers/global"
	"github.com/labstack/gommon/log"
)

func Logger() {
	c := global.GIANT_CONF.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := log.New(c.Prefix)
	global.GIANT_LOG = logger
}
