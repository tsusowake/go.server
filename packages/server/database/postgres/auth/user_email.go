package auth

import (
	"context"

	"github.com/tsusowake/go.server/database/generated"
	"github.com/tsusowake/go.server/domain/auth/model"
)

type userEmail struct {
	query *generated.Queries
}

var _ UserEmail = (*userEmail)(nil)

func NewUserEmail(q *generated.Queries) UserEmail {
	return &userEmail{
		query: q,
	}
}

func (u *userEmail) GetByUserID(
	ctx context.Context,
	userID string,
) (*model.UserEmail, error) {
	ret, e := u.query.GetByUserID(ctx, userID)
	return &model.UserEmail{
		UserID: ret.UserID,
		Email:  ret.Email,
	}, e
}
