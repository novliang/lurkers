package initialization

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/novliang/lurkers/global"
)

const defaultConfigFile = "config.yaml"

func Config() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.GIANT_CONF); err != nil {
		fmt.Println(err)
	}
	global.GIANT_VIPER = v
}
