package setting

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Rdb *redis.Client

func InitRedis() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return err
}
