package database

//go:generate mockgen -source=./database.go -package=mock -destination=./mock/mock.go

import (
	"context"

	"github.com/tsusowake/go.server/internal/domain/auth/entity"
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
type UserEmail interface {
	GetUserEmailByUserID(ctx context.Context, userID string) (*entity.UserEmail, error)
}
type UserLock interface{}
