package database

import (
	"gorm.io/gorm"

	"github.com/tsusowake/go.server/database/postgres/auth"
)

type Database struct {
	Auth *auth.Repository
	// Membership
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		Auth: auth.NewRepository(db),
	}
}
