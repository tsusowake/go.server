package auth

//go:generate mockgen -source=./repository.go -package=mock -destination=./mock/mock.go

import (
	"context"

	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/auth/entity"
)

type Repository struct {
	User           User
	UserEmail      UserEmail
	UserCredential UserCredential
	UserLock       UserLock
}

func NewRepository(q *generated.Queries) *Repository {
	return &Repository{
		User:           NewUser(q),
		UserEmail:      NewUserEmail(q),
		UserCredential: nil,
		UserLock:       nil,
	}
}

type User interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context) (string, error)
}

type UserCredential interface{}
type UserEmail interface {
	GetByUserID(ctx context.Context, userID string) (*entity.UserEmail, error)
}
type UserLock interface{}
