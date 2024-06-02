package auth

import (
	"context"

	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/auth/entity"
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
) (*entity.UserEmail, error) {
	ret, e := u.query.GetByUserID(ctx, userID)
	return &entity.UserEmail{
		UserID: ret.UserID,
		Email:  ret.Email,
	}, e
}
