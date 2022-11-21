package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisDemo() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("score:", val)

	val2, err := rdb.Get("keytest").Result()
	if err == redis.Nil {
		// 这里主要是看这个key不存在的判定方法就可以了
		fmt.Println("keytest does not exist")
	} else if err != nil {
		fmt.Println("get keytest failed")
	} else {
		fmt.Println("keytest:", val2)
	}

}
