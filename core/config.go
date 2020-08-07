package core

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/ibyond/go-start/global"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.GstConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GstConfig); err != nil {
		fmt.Println(err)
	}
	global.GstVp = v
}
