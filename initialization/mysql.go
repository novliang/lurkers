package initialization

import (
	"github.com/novliang/lurkers/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Mysql() {
	admin := global.GIANT_CONF.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.GIANT_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.GIANT_DB = db
		global.GIANT_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GIANT_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GIANT_DB.LogMode(admin.LogMode)
	}
}
