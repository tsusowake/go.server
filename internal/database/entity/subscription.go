package entity

import (
	"time"
)

type Subscription struct {
	ID                     string    `db:"id"`
	UserID                 string    `db:"user_id"`
	ProductID              string    `db:"product_id"`
	PlatformSubscriptionID string    `db:"platform_subscription_id"`
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
}
