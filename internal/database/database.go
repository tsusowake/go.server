package database

//go:generate mockgen -source=./database.go -package=mock -destination=./mock/mock.go

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/tsusowake/go.server/internal/database/entity"
)

type Database struct {
	User        User
	UserSetting UserSetting
}

type User interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type UserSetting interface {
	CreateTx(ctx context.Context, tx *sqlx.Tx, setting *entity.UserSetting) error
}

// TODO contracts
