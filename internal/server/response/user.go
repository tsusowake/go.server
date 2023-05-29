package response

import (
	"github.com/tsusowake/go.server/internal/database/entity"
	"github.com/tsusowake/go.server/pkg/conv"
	"time"
)

type User struct {
	ID         string            `json:"id"`
	Password   string            `json:"password"`
	Email      string            `json:"email"`
	Status     entity.UserStatus `json:"status"`
	DisabledAt *time.Time        `json:"disabled_at,omitempty"`
	BannedAt   *time.Time        `json:"banned_at,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

func ToUser(u *entity.User) *User {
	return &User{
		ID:         u.ID,
		Password:   u.Password,
		Email:      u.Email,
		Status:     u.Status,
		DisabledAt: conv.NullTimeToTime(u.DisabledAt),
		BannedAt:   conv.NullTimeToTime(u.BannedAt),
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
