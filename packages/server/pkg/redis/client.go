package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Close() error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Do(ctx context.Context, args ...any) error
	Publish(ctx context.Context, channel string, message any) error
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}

type redisClient struct {
	rdb *redis.Client
}

func (rc *redisClient) Close() error {
	return rc.rdb.Close()
}

func (rc *redisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return rc.rdb.Set(ctx, key, value, expiration).Err()
}

func (rc *redisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}

func (rc *redisClient) Do(ctx context.Context, args ...any) error {
	return rc.rdb.Do(ctx, args...).Err()
}

func (rc *redisClient) Publish(ctx context.Context, channel string, message any) error {
	return rc.rdb.Publish(ctx, channel, message).Err()
}

func (rc *redisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return rc.rdb.Subscribe(ctx, channels...)
}

func NewRedisClient(
	ctx context.Context,
	host string,
	password string,
) RedisClient {
	// TODO use Conn, close
	rdb := redis.NewClient(&redis.Options{
		Addr:     host, // "localhost:6379",
		Password: password,
		DB:       0,
	})
	return &redisClient{
		rdb: rdb,
	}
}
