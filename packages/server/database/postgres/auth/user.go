package auth

import (
	"context"

	"gorm.io/gorm"

	"github.com/tsusowake/go.server/database/entity"
	"github.com/tsusowake/go.server/domain/auth/model"
)

type user struct {
	db *gorm.DB
}

var _ User = (*user)(nil)

func NewUser(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (u *user) GetByID(ctx context.Context, id string) (*model.User, error) {
	var ret entity.User
	if err := u.db.WithContext(ctx).First(&ret, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return u.toEntity(ctx, &ret), nil
}

func (u *user) Create(ctx context.Context) (string, error) {
	var id string
	err := u.db.WithContext(ctx).
		Raw("insert into users default values returning id").
		Scan(&id).Error
	return id, err
}

func (u *user) toEntity(_ context.Context, uu *entity.User) *model.User {
	return &model.User{
		ID: uu.ID,
	}
}
