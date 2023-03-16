package myclient

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

// 初始化连接
func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func hset(hash, key, value string) error {

	if rdb == nil {
		if err := initRedisClient(); err != nil {
			return err
		}
	}

	ctx := context.Background()

	_, err := rdb.HSet(ctx, hash, key, value).Result()
	if err != nil {
		return fmt.Errorf("error while doing HSET command in gredis : %v", err)
	}

	return err
}

func hget(hash, key string) (string, error) {

	if rdb == nil {
		if err := initRedisClient(); err != nil {
			return "", err
		}
	}

	ctx := context.Background()

	value, err := rdb.HGet(ctx, hash, key).Result()
	if err != nil {
		return value, fmt.Errorf("error while doing HGET command in gredis : %v", err)
	}

	return value, err

}
