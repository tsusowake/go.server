package auth

import (
	"context"

	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/auth/entity"
)

type user struct {
	query *generated.Queries
}

var _ User = (*user)(nil)

func NewUser(q *generated.Queries) User {
	return &user{
		query: q,
	}
}

func (u *user) GetByID(ctx context.Context, id string) (*entity.User, error) {
	ret, e := u.query.GetByID(ctx, id)
	return u.toEntity(ctx, &ret), e
}

func (u *user) Create(ctx context.Context) (string, error) {
	id, err := u.query.Create(ctx)
	return id, err
}

func (u *user) toEntity(_ context.Context, uu *generated.User) *entity.User {
	return &entity.User{
		ID: uu.ID,
	}
}
