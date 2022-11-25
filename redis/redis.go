package redis

import (
	"fmt"
	"time"

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

func RedisDemo2() {
	zsetkey := "language_rank"
	languages := []redis.Z{
		{Score: 90, Member: "java"},
		{Score: 80, Member: "go"},
		{Score: 70, Member: "js"},
		{Score: 60, Member: "rust"},
		{Score: 50, Member: "c++"},
	}

	num, err := rdb.ZAdd(zsetkey, languages...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("num:", num)
	// 增加值
	newScore, err := rdb.ZIncrBy(zsetkey, 10, "go").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("newScore:", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(zsetkey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println("name:", z.Member,
			"  score:", z.Score)
	}

	// 取分数在一定范围之内的
	op := redis.ZRangeBy{
		Min: "80",
		Max: "110",
	}

	ret = rdb.ZRangeByScoreWithScores(zsetkey, op).Val()
	for _, z := range ret {
		fmt.Println(z.Member, "  ", z.Score)
	}

}

func WatchDemo() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int64()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.TxPipelined(func(pipeliner redis.Pipeliner) error {
			time.Sleep(20 * time.Second)
			pipeliner.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)

	if err != nil {
		fmt.Println("tx exec failed:", err)
		return
	}
	fmt.Println("tx exec success")
}
