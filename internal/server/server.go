package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tsusowake/go.server/internal/pkg/logger"
	"github.com/tsusowake/go.server/internal/pkg/redis"
)

type server struct {
	EchoServer *echo.Echo
	Logger     *zap.Logger

	RedisClient redis.RedisClient
}

func NewServer(ctx context.Context) (*server, error) {
	e := echo.New()

	logger, err := logger.NewLogger(zapcore.DebugLevel)
	if err != nil {
		return nil, err
	}
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:   true,
		LogProtocol:  true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogRequestID: true,
		LogReferer:   true,
		LogUserAgent: true,
		LogStatus:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = errorHandler

	rc := redis.NewRedisClient(ctx, "localhost:6379", "")

	s := &server{
		EchoServer:  e,
		Logger:      logger,
		RedisClient: rc,
	}
	s.setupHandlers(e)
	return s, nil
}

func (s *server) Start(port string) error {
	return s.EchoServer.Start(fmt.Sprintf(":%s", port))
}

func (s *server) Stop() error {
	if err := s.RedisClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *server) setupHandlers(e *echo.Echo) {
	// /rooms
	e.POST("/rooms", s.createRoom)
	// TODO channel の subscribe は別サーバーが良いかも
	e.GET("/rooms/chat", s.connectChat)
	e.POST("/rooms/chat/messages", s.sendMessage)
}

func errorHandler(err error, ectx echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	ectx.Logger().Error("InternalServerError: ",
		zap.Int("code", code),
		zap.String("cause", err.Error()),
	)
	// errorPage := fmt.Sprintf("%d.html", code)
	// if err := c.File(errorPage); err != nil {
	// 	c.Logger().Error(err)
	// }
}
