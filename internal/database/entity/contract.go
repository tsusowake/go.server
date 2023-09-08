package entity

import (
	"time"
)

type Contract struct {
	ID             string    `db:"id"`
	UserID         string    `db:"user_id"`
	PlanID         string    `db:"plan_id"`
	SubscriptionID string    `db:"subscription_id"`
	StartedAt      time.Time `db:"started_at"`
	EndedAt        time.Time `db:"ended_at"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
