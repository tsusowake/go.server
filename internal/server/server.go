package server

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	db "github.com/tsusowake/go.server/internal/database"
	mysql2 "github.com/tsusowake/go.server/internal/database/mysql"
	"github.com/tsusowake/go.server/pkg/database"
	"github.com/tsusowake/go.server/pkg/logger"
	"github.com/tsusowake/go.server/pkg/redis"
	ptime "github.com/tsusowake/go.server/pkg/time"
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

	logger, err := logger.NewLogger(zapcore.DebugLevel)
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
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
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
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

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
	db, err := database.Open(mconf, logger)
	if err != nil {
		return err
	}
	defer db.DB.Close()
	// See "Important settings" section.
	db.DB.SetConnMaxLifetime(time.Minute * 3)
	db.DB.SetMaxOpenConns(10)
	db.DB.SetMaxIdleConns(10)

	s := &server{
		EchoServer:  e,
		Logger:      logger,
		RedisClient: rc,
		Database:    mysql2.NewDatabase(db),
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
