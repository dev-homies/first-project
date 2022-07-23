package core

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func SetupConfig() {
	Config = viper.New()
	Config.AutomaticEnv()

	Config.SetConfigFile(".env")
	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}

	Config.WatchConfig()
	Config.OnConfigChange(func(event fsnotify.Event) {
		Logger.Infof("Config file changed: %s", event.Name)
	})
}
