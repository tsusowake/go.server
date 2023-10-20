package server

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	db "github.com/tsusowake/go.server/internal/database"
	mysql2 "github.com/tsusowake/go.server/internal/database/mysql"
	"github.com/tsusowake/go.server/util/database"
	"github.com/tsusowake/go.server/util/echoutil"
	"github.com/tsusowake/go.server/util/logger"
	"github.com/tsusowake/go.server/util/redis"
	ptime "github.com/tsusowake/go.server/util/time"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type server struct {
	EchoServer  *echo.Echo
	Logger      *zap.Logger
	RedisClient redis.RedisClient
	Database    *db.Database
}

func Run(ctx context.Context) error {
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
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &
		}
	})
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

	loc, err := ptime.LoadLocation()
	if err != nil {
		return err
	}
	mconf := &mysql.Config{
		User:      "test",
		Passwd:    "test",
		Addr:      "127.0.0.1",
		DBName:    "yunne_test",
		Collation: "utf8_general_ci",
		Loc:       loc,
	}
	mconf = database.WithParseTime(mconf)
	dbc, err := database.Open(mconf, l)
	if err != nil {
		return err
	}
	defer func() {
		_ = dbc.DB.Close()
	}()
	// See "Important settings" section.
	dbc.DB.SetConnMaxLifetime(time.Minute * 3)
	dbc.DB.SetMaxOpenConns(10)
	dbc.DB.SetMaxIdleConns(10)

	s := &server{
		EchoServer:  e,
		Logger:      l,
		RedisClient: rc,
		Database:    mysql2.NewDatabase(dbc),
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
	// TODO 副作用を消す
	e.GET("/user/create", s.createUser)

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
