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
	"github.com/labstack/echo/v4"
	"github.com/morikuni/failure/v2"
	"golang.org/x/sync/errgroup"

	"github.com/tsusowake/go.server/config"
	"github.com/tsusowake/go.server/database"
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

	// Start server
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		go func() {
			<-ctx.Done()
			if stopErr := stopServer(ctx, e); stopErr != nil {
				slog.ErrorContext(ctx, "failed to shutdown echo server", slog.String("error", err.Error()))
			}
		}()
		return startServer(egCtx, e)
	})

	if egErr := eg.Wait(); egErr != nil && !errors.Is(egErr, context.Canceled) {
		return failure.Wrap(egErr, failure.Message("server encountered an error"))
	}
	return nil
}

func startServer(ctx context.Context, e *echo.Echo) error {
	slog.InfoContext(ctx, "start server")
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return failure.Wrap(err, failure.Message("failed to start server"))
	}
	return nil
}

func stopServer(ctx context.Context, e *echo.Echo) error {
	// graceful shutdown
	slog.InfoContext(ctx, "stop server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return failure.Wrap(err, failure.Message("failed to shutdown server"))
	}
	slog.InfoContext(ctx, "server stopped")
	return nil
}
