package database

import (
	"errors"
	"github.com/cenkalti/backoff/v4"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Connector struct {
	DB     *sqlx.DB
	Logger *zap.Logger
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
