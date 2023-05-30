package mysql

import (
	"github.com/tsusowake/go.server/internal/database"
	db "github.com/tsusowake/go.server/pkg/database"
)

func NewDatabase(conn *db.Connector) *database.Database {
	return &database.Database{
		User:        &user{conn: conn},
		UserSetting: &userSetting{conn: conn},
	}
}
