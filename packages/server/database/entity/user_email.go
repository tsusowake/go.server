package entity

import "time"

// UserEmail is the GORM entity for the "user_emails" table.
type UserEmail struct {
	UserID    string `gorm:"column:user_id;primaryKey"`
	Email     string `gorm:"column:email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName returns the table name for the UserEmail entity.
func (UserEmail) TableName() string {
	return "user_emails"
}
