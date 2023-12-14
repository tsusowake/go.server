package database

//go:generate mockgen -source=./database.go -package=mock -destination=./mock/mock.go

import (
	"context"

	"github.com/tsusowake/go.server/internal/domain/account/entity"
)

type Database struct {
	User           User
	UserEmail      UserEmail
	UserCredential UserCredential
	UserLock       UserLock
}

type User interface {
	GetByID(ctx context.Context, id uint64) (*entity.User, error)
	Create(ctx context.Context) (uint64, error)
}

type UserCredential interface{}
type UserEmail interface{}
type UserLock interface{}
