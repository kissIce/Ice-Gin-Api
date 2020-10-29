package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"ice/global"
)

const ConfigFile = "./config/config.yaml"

func init()  {
	var config string
	flag.StringVar(&config, "c", "", "choose config file")
	flag.Parse()
	if config == "" {
		config = ConfigFile
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail load config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.IceConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.IceConfig); err != nil {
		fmt.Println(err)
	}
	global.IceVp = v
}