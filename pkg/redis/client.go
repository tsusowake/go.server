package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type redisClient struct {
	rdb *redis.Client
}

func (rc *redisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return rc.rdb.Set(ctx, key, value, expiration).Err()
}

func (rc *redisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}

func NewRedisClient(
	ctx context.Context,
	host string,
	password string,
) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host, // "localhost:6379",
		Password: password,
		DB:       0,
	})

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist

	return &redisClient{
		rdb: rdb,
	}
}
