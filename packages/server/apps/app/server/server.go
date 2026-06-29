package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/morikuni/failure/v2"

	"github.com/tsusowake/go.server/apps/app/server/handlers"
	"github.com/tsusowake/go.server/config"
	"github.com/tsusowake/go.server/database"
	"github.com/tsusowake/go.server/domain"
	"github.com/tsusowake/go.server/database/generated"
	"github.com/tsusowake/go.server/pkg/logger"
	pkgmiddleware "github.com/tsusowake/go.server/pkg/middleware"
	"github.com/tsusowake/go.server/pkg/redis"
	pkgtime "github.com/tsusowake/go.server/pkg/time"
	"github.com/tsusowake/go.server/pkg/ulid"
)

var conf config.Config

type server struct {
	EchoServer    *echo.Echo
	RedisClient   redis.RedisClient
	Database      *database.Database
	Clocker       pkgtime.Clocker
	ULIDGenerator ulid.ULIDGenerator
}

func Run(ctx context.Context) error {
	conf = env.Must(env.ParseAs[config.Config]())

	l := logger.NewLogger(os.Stdout, slog.LevelInfo, logger.LogMetadata{
		Service: conf.Service,
		Env:     conf.Env,
	})
	slog.SetDefault(l)

	if err := runServer(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run server", slog.String("error", err.Error()))
		return failure.Wrap(err)
	}
	return nil
}

func runServer(ctx context.Context) error {
	// RDB
	dbPoolCtx, cancelDbPoolCtx := context.WithCancel(ctx)
	defer cancelDbPoolCtx()
	dbpool, err := pgxpool.New(dbPoolCtx, conf.DBConfig.ConnString())
	if err != nil {
		return failure.Wrap(err, failure.Message("failed to connect to database"))
	}
	defer dbpool.Close()

	// Redis
	redisCtx, cancelRedisCtx := context.WithCancel(ctx)
	defer cancelRedisCtx()
	rc := redis.NewRedisClient(redisCtx, conf.RedisConfig.Address(), conf.RedisConfig.Password)
	defer func() {
		if err := rc.Close(); err != nil {
			slog.ErrorContext(ctx, "failed to close redis client", slog.String("error", err.Error()))
		}
	}()

	// Setup Echo server
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
		Clocker:       pkgtime.NewClocker(),
		ULIDGenerator: ulid.NewULIDGenerator(),
	}
	fmt.Println("server initialized: ", s)

	// Register OpenAPI-generated routes (strict handlers).
	base := handlers.NewBaseHandler(domain.NewRepository())
	handlers.NewAPI(base).Register(e)

	// Start server. StartConfig.Start blocks until ctx is cancelled, then
	// gracefully shuts the server down within GracefulTimeout.
	slog.InfoContext(ctx, "start server")
	sc := echo.StartConfig{
		Address:         ":8080",
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, e); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return failure.Wrap(err, failure.Message("server encountered an error"))
	}
	slog.InfoContext(ctx, "server stopped")
	return nil
}
