package redis

import (
	"fmt"
	"hjfu/Wolverine/domain"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Close() {
	_ = rdb.Close()
}

func Init(config *domain.RedisConfig) (err error) {
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
