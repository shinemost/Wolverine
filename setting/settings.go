package setting

import (
	"fmt"
	"hjfu/Wolverine/domain"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(domain.AppConfig)

func InitConfig(filePath string) error {

	// 如果设置filepath就直接使用 否则用当前目录的
	if len(filePath) == 0 {
		viper.AddConfigPath(".")

		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
	} else {
		viper.SetConfigFile(filePath)
	}

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("viper init failed:", err)
		return err
	}

	if err := viper.Unmarshal(Config); err != nil {
		fmt.Println("viper Unmarshal err", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
	})

	return err

}
