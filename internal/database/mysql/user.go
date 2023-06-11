package mysql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/tsusowake/go.server/internal/database/entity"
	"github.com/tsusowake/go.server/util/database"
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
	conn        *database.Connector
	userSetting *userSetting
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

func (u *user) Create(ctx context.Context, user *entity.User) error {
	id, err := u.conn.UUID()
	if err != nil {
		return err
	}
	user.ID = id
	now := u.conn.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err = database.DoTx(ctx, u.conn, func(txCtx context.Context, tx *sqlx.Tx) (any, error) {
		sql, args, err := sq.Insert(userTable).
			Columns("id", "password", "email", "status", "created_at", "updated_at").
			Values(
				user.ID,
				user.Password,
				user.Email,
				user.Status,
				user.CreatedAt,
				user.UpdatedAt,
			).
			ToSql()
		if err != nil {
			return nil, err
		}
		if _, err := tx.ExecContext(txCtx, sql, args...); err != nil {
			return nil, err
		}
		setting := &entity.UserSetting{
			UserID:    user.ID,
			Language:  entity.LanguageJa,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := u.userSetting.CreateTx(txCtx, tx, setting); err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
