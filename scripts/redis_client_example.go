package main

import (
	"context"
	"fmt"
	"github.com/tsusowake/go.server/pkg/redis"
)

func main() {
	ctx := context.Background()
	rc := redis.NewRedisClient(ctx, "localhost:6379", "")

	if err := rc.Set(ctx, "key.1", "test.value", 0); err != nil {
		panic(fmt.Errorf("set error: %s", err))
	}
	if val, err := rc.Get(ctx, "key.1"); err != nil {
		panic(fmt.Errorf("get error: %s", err))
	} else {
		fmt.Printf("get value: %s\n", val)
	}
}
