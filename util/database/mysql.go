package database

import (
	"context"
	"errors"
	"github.com/cenkalti/backoff/v4"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tsusowake/go.server/util/uuid"
	"go.uber.org/zap"
	"time"
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

func (c *Connector) Begin(
	ctx context.Context,
	txFunc func(ctx context.Context, tx *sqlx.Tx) (any, error),
) (any, error) {
	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			c.Rollback(tx)
			panic(r)
		}
	}()
	ret, err := txFunc(ctx, tx)
	if err != nil {
		c.Rollback(tx)
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ret, nil
}
