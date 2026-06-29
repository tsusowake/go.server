package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_redisClient_Publish(t *testing.T) {
	ctx := context.Background()
	cli := NewRedisClient(ctx, "localhost:6379", "")

	t.Run("success", func(t *testing.T) {
		err := cli.Publish(ctx, "test.channel", "test.message")
		require.NoError(t, err)
	})
}
