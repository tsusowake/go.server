package repository

import (
	"context"

	"github.com/tsusowake/go.server/internal/database"
	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/account/entity"
)

type User struct {
	query *generated.Queries
}

var _ database.User = (*User)(nil)

func NewUser(q *generated.Queries) *User {
	return &User{
		query: q,
	}
}

func (u *User) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	ret, e := u.query.GetUserByID(ctx, id)
	return u.ToEntity(ctx, &ret), e
}

func (u *User) Create(ctx context.Context) (uint64, error) {
	return u.query.CreateUser(ctx)
}

func (u *User) ToEntity(_ context.Context, uu *generated.User) *entity.User {
	return &entity.User{
		ID: uu.ID,
	}
}
