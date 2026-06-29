package entity

import "time"

// User is the GORM entity for the "users" table.
type User struct {
	ID        string `gorm:"column:id;primaryKey"`
	CreatedAt time.Time
}

// TableName returns the table name for the User entity.
func (User) TableName() string {
	return "users"
}
