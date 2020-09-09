package global

import (
	"github.com/novliang/lurkers/config"
	"github.com/novliang/lurkers/storage"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var (
	GIANT_DB            *gorm.DB
	GIANT_CONF          config.Server
	GIANT_LOG           *log.Logger
	GIANT_VIPER         *viper.Viper
	GIANT_STORAGE       *storage.Storage
)
