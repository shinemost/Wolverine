package setting

import (
	"fmt"
	"hjfu/Wolverine/domain"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func CloseRedis() {
	_ = rdb.Close()
}

func InitRedis(config *domain.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Post),
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return err
}
