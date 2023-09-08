package entity

import (
	"time"
)

type Platform uint8

const (
	PlatformNone Platform = iota
	PlatformWeb
	PlatformIOS
	PlatformAndroid
)

type PlanPrice struct {
	ID        string    `db:"id"`
	PlanID    string    `db:"plan_id"`
	ProductID string    `db:"product_id"`
	Platform  Platform  `db:"platform"`
	Price     uint32    `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
