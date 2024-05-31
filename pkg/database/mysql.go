package database

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/tsusowake/go.server/pkg/uuid"
)

type Connector struct {
	DB     *sqlx.DB
	Logger *zap.Logger
	Now    func() time.Time
	UUID   func() (string, error)
}

func Open(conf *mysql.Config, logger *zap.Logger) (*Connector, error) {
	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = ping(db, 10); err != nil {
		return nil, err
	}
	conn := &Connector{
		DB:     db,
		Logger: logger,
		Now:    time.Now,
		UUID:   uuid.NewURLSafeString,
	}
	return conn, nil
}

func WithParseTime(conf *mysql.Config) *mysql.Config {
	params := conf.Params
	if len(params) == 0 {
		params = map[string]string{}
	}
	params["parseTime"] = "true"
	conf.Params = params
	return conf
}

func ping(db *sqlx.DB, count int) (err error) {
	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(count))
	if err := backoff.Retry(db.Ping, b); err != nil {
		return errors.New("mysql: Failed to connect to database")
	}
	return nil
}

func (c *Connector) Rollback(tx *sqlx.Tx) {
	if err := tx.Rollback(); err != nil {
		c.Logger.Error("failed to rollback", zap.Error(err))
	}
}

// DoTx
// see https://github.com/golang/go/issues/49085
func DoTx[T any](
	ctx context.Context,
	c *Connector,
	fn func(ctx context.Context, tx *sqlx.Tx) (T, error),
) (ret T, err error) {
	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		return ret, err
	}
	defer func() {
		if r := recover(); r != nil {
			c.Rollback(tx)
			panic(r)
		}
	}()
	ret, err = fn(ctx, tx)
	if err != nil {
		c.Rollback(tx)
		return ret, err
	}
	if err = tx.Commit(); err != nil {
		return ret, err
	}
	return ret, nil
}
