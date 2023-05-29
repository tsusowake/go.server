package entity

import (
	"database/sql"
	"time"
)

type UserStatus uint8

const (
	UserStatusNone UserStatus = iota
	UserStatusActive
	UserStatusDisabled
	UserStatusBanned
)

type User struct {
	ID         string       `db:"id"`
	Password   string       `db:"password"`
	Email      string       `db:"email"`
	Status     UserStatus   `db:"status"`
	DisabledAt sql.NullTime `db:"disabled_at"`
	BannedAt   sql.NullTime `db:"banned_at"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
}
