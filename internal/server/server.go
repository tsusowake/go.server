package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/tsusowake/go.server/internal/config"
	"github.com/tsusowake/go.server/internal/database"
	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/pkg/logger"
	pkgmiddleware "github.com/tsusowake/go.server/pkg/middleware"
	"github.com/tsusowake/go.server/pkg/redis"
	"github.com/tsusowake/go.server/pkg/time"
	"github.com/tsusowake/go.server/pkg/ulid"
)

var conf config.Config

type server struct {
	EchoServer    *echo.Echo
	RedisClient   redis.RedisClient
	Database      *database.Database
	Clocker       time.Clocker
	ULIDGenerator ulid.ULIDGenerator
}

func Run(ctx context.Context) error {
	conf = env.Must(env.ParseAs[config.Config]())

	fd := os.Stdout
	l := logger.NewLogger(fd, slog.LevelInfo, logger.LogMetadata{
		Service: conf.Service,
		Env:     conf.Env,
	})
	slog.SetDefault(l)

	if err := runServer(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run server", slog.String("error", err.Error()))
		return err
	}
	return nil
}

func runServer(ctx context.Context) error {
	// RDB
	dbPoolCtx, cancelDbPoolCtx := context.WithCancel(ctx)
	defer cancelDbPoolCtx()
	dbpool, err := pgxpool.New(dbPoolCtx, conf.DBConfig.ConnString())
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer dbpool.Close()

	// Redis
	redisCtx, cancelRedisCtx := context.WithCancel(ctx)
	defer cancelRedisCtx()
	rc := redis.NewRedisClient(redisCtx, conf.RedisConfig.Address(), conf.RedisConfig.Password)

	e := echo.New()

	e.HTTPErrorHandler = pkgmiddleware.ErrorHandler
	e.Validator = pkgmiddleware.NewCustomValidator()

	// setup middlewares
	e.Use(pkgmiddleware.CORSMiddleware(conf.CORSConfig.Decode()))
	e.Use(pkgmiddleware.NewLoggerMiddleware())
	e.Use(pkgmiddleware.RecoverMiddleware())

	query := generated.New(dbpool)

	s := &server{
		EchoServer:    e,
		RedisClient:   rc,
		Database:      database.NewDatabase(query),
		Clocker:       time.NewClocker(),
		ULIDGenerator: ulid.NewULIDGenerator(),
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
	// API サーバーとして実行するので /favicon.ico は不要だがエラーログが邪魔
	e.GET("/favicon.ico", s.getFavicon)

	// users
	e.GET("/user/:id", s.getUser)
	e.POST("/user", s.createUser)
}
