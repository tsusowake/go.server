package postgres

import (
	"github.com/tsusowake/go.server/internal/database"
	"github.com/tsusowake/go.server/internal/database/generated"
	"github.com/tsusowake/go.server/internal/domain/account/repository"
)

func NewDatabase(q *generated.Queries) *database.Database {
	return &database.Database{
		// account
		User:           repository.NewUser(q),
		UserEmail:      nil,
		UserCredential: nil,
		UserLock:       nil,

		// membership
	}
}
