package repository

import (
	"context"

	"github.com/tsusowake/go.server/internal/database"
	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/auth/entity"
)

type UserEmail struct {
	query *generated.Queries
}

var _ database.UserEmail = (*UserEmail)(nil)

func NewUserEmail(q *generated.UserEmail) *UserEmail {
	return &UserEmail{
		query: q,
	}
}

func (u *UserEmail) GetByUserID(
	ctx context.Context,
	userID string,
) (*entity.UserEmail, error) {
	ret, e := u.query.GetUserEmailByUserID(ctx, userID)
	return &entity.UserEmail{
		UserID: ret.UserID,
		Email:  ret.Email,
	}, e
}
