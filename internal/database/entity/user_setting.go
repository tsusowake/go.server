package entity

import (
	"time"
)

type Language string

const (
	LanguageJa   Language = "ja"
	LanguageEnUS Language = "en-US"
	LanguageEnGB Language = "en-GB"
	LanguageEnCA Language = "en-CA"
	LanguageEnAU Language = "en-AU"
	// ...
)

type UserSetting struct {
	UserID    string    `db:"user_id"`
	Language  Language  `db:"language"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
