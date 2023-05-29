package database

//go:generate mockgen -source=./database.go -package=mock -destination=./mock/mock.go

import (
	"context"
	"github.com/tsusowake/go.server/internal/database/entity"
)

type Database struct {
	User User
}

type User interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
