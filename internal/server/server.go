package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tsusowake/go.server/internal/config"
	"github.com/tsusowake/go.server/internal/database"
	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/database/postgres"
	"github.com/tsusowake/go.server/util/echoutil"
	"github.com/tsusowake/go.server/util/logger"
	"github.com/tsusowake/go.server/util/redis"
)

type server struct {
	EchoServer  *echo.Echo
	Logger      *zap.Logger
	RedisClient redis.RedisClient
	Database    *database.Database
}

func Run(ctx context.Context) error {
	var c config.Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return err
	}

	e := echo.New()

	l, err := logger.NewLogger(zapcore.DebugLevel)
	if err != nil {
		return err
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
			l.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		//cc := &
	//	}
	//})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			// TODO
		},
		// AllowHeaders: []string{echoutil.HeaderOrigin, echoutil.HeaderContentType, echoutil.HeaderAccept},
		AllowMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = errorHandler
	echoutil.UseCustomValidator(e)

	rc := redis.NewRedisClient(ctx, "localhost:6379", "")

	// RDB
	dbPoolCtx, cancelDbPoolCtx := context.WithCancel(ctx)
	defer cancelDbPoolCtx()
	dbpool, err := pgxpool.New(dbPoolCtx, c.DB.ConnString())
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer dbpool.Close()

	query := generated.New(dbpool)

	s := &server{
		EchoServer:  e,
		Logger:      l,
		RedisClient: rc,
		Database:    postgres.NewDatabase(query),
	}
	s.setupHandlers(e)
	return s.Start("1323")
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
	// API サーバーとして実行するので /favicon.ico はリクエストされない
	e.GET("/favicon.ico", s.getFavicon)

	// /users
	e.GET("/user/:id", s.getUser)

	// /rooms
	e.POST("/rooms", s.createRoom)
	// TODO channel の subscribe は別サーバーが良いかも
	e.GET("/rooms/chat", s.connectChat)
	e.POST("/rooms/chat/messages", s.sendMessage)
}

func errorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	ctx.Logger().Error("InternalServerError: ",
		zap.Int("code", code),
		zap.String("cause", err.Error()),
	)
}
