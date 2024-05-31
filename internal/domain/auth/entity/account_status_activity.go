package entity

import "time"

type AccountStatusActivity struct {
	ID           string
	UserID       string
	ActivityType int16
	OccurredAt   time.Time
}
