package auth

//go:generate mockgen -source=./repository.go -package=mock -destination=./mock/mock.go

import (
	"context"

	"gorm.io/gorm"

	"github.com/tsusowake/go.server/domain/auth/model"
)

type Repository struct {
	User           User
	UserEmail      UserEmail
	UserCredential UserCredential
	UserLock       UserLock
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:           NewUser(db),
		UserEmail:      NewUserEmail(db),
		UserCredential: nil,
		UserLock:       nil,
	}
}

type User interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context) (string, error)
}

type UserCredential interface{}
type UserEmail interface {
	GetByUserID(ctx context.Context, userID string) (*model.UserEmail, error)
}
type UserLock interface{}
