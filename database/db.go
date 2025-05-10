package database

import (
	"github.com/tsusowake/go.server/database/generated"
	"github.com/tsusowake/go.server/database/postgres/auth"
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
