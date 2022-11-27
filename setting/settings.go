package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() error {

	viper.AddConfigPath(".")

	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("viper init failed:", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
	})

	return err

}
