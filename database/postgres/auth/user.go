package auth

import (
	"context"

	"github.com/tsusowake/go.server/database/generated"
	"github.com/tsusowake/go.server/domain/auth/model"
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

func (u *user) GetByID(ctx context.Context, id string) (*model.User, error) {
	ret, e := u.query.GetByID(ctx, id)
	return u.toEntity(ctx, &ret), e
}

func (u *user) Create(ctx context.Context) (string, error) {
	id, err := u.query.Create(ctx)
	return id, err
}

func (u *user) toEntity(_ context.Context, uu *generated.User) *model.User {
	return &model.User{
		ID: uu.ID,
	}
}
