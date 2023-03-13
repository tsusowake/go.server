package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// TODO sse
func (s *server) connectChat(ectx echo.Context) error {
	ctx := FromEchoContext(ectx)

	writer := ectx.Response().Writer
	flusher, ok := writer.(http.Flusher)
	if !ok {
		return errors.New("flusher: could not cast to http.Flusher")
	}
	req := ectx.Request()

	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")

	sub := s.RedisClient.Subscribe(ctx, "test.channel.1")
	iface, err := sub.Receive(ctx)
	if err != nil {
		return fmt.Errorf("subscribe error, %s", err)
	}

	switch iface.(type) {
	case *redis.Subscription:
		// subscribe succeeded
		s.Logger.Info("connect: room")
	// case *redis.Message:
	// 	// received first message
	// case *redis.Pong:
	// 	// pong received
	default:
		return errors.New("failed to subscribe channel")
	}

	ch := sub.Channel()
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Fprintf(writer, "message: %s\n", msg.String())
				flusher.Flush()
			}
		}
	}()
	<-req.Context().Done()
	return nil
}

func (s *server) createRoom(ectx echo.Context) error {
	ctx := FromEchoContext(ectx)
	if err := s.RedisClient.Publish(ctx, "test.channel.1", "test.publish"); err != nil {
		return err
	}
	// TODO
	return nil
}

func (s *server) sendMessage(ectx echo.Context) error {
	ctx := FromEchoContext(ectx)
	msg := fmt.Sprintf("message.now:%s", time.Now().Format(time.RFC3339))
	if err := s.RedisClient.Publish(ctx, "test.channel.1", msg); err != nil {
		return err
	}
	return nil
}
