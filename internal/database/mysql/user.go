package mysql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/tsusowake/go.server/internal/database/entity"
	"github.com/tsusowake/go.server/pkg/database"
)

const (
	userTable = "user"
)

var userColumns = []string{
	"id",
	"password",
	"email",
	"status",
	"disabled_at",
	"banned_at",
	"created_at",
	"updated_at",
}

type user struct {
	conn *database.Connector
}

func (u *user) GetByID(ctx context.Context, id string) (*entity.User, error) {
	sql, _, err := sq.Select(userColumns...).
		From(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var dest entity.User
	if err := u.conn.DB.GetContext(ctx, &dest, sql, id); err != nil {
		return nil, err
	}
	return &dest, nil
}
