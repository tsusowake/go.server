package auth

import (
	"context"

	"gorm.io/gorm"

	"github.com/tsusowake/go.server/database/entity"
	"github.com/tsusowake/go.server/domain/auth/model"
)

type userEmail struct {
	db *gorm.DB
}

var _ UserEmail = (*userEmail)(nil)

func NewUserEmail(db *gorm.DB) UserEmail {
	return &userEmail{
		db: db,
	}
}

func (u *userEmail) GetByUserID(
	ctx context.Context,
	userID string,
) (*model.UserEmail, error) {
	var ret entity.UserEmail
	if err := u.db.WithContext(ctx).First(&ret, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &model.UserEmail{
		UserID: ret.UserID,
		Email:  ret.Email,
	}, nil
}
