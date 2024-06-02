package database

import (
	"github.com/tsusowake/go.server/internal/database/postgres/auth"

	"github.com/tsusowake/go.server/internal/database/generated"
)

type Database struct {
	Auth *auth.Repository
	// Membership
}

func NewDatabase(q *generated.Queries) *Database {
	return &Database{
		Auth: auth.NewRepository(q),
	}
}
