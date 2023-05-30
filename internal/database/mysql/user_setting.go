package mysql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/tsusowake/go.server/internal/database/entity"
	"github.com/tsusowake/go.server/pkg/database"
)

const userSettingTable = "user_setting"

var userSettingColumns = []string{
	"user_id",
	"language",
	"created_at",
	"updated_at",
}

type userSetting struct {
	conn *database.Connector
}

func (us *userSetting) CreateTx(ctx context.Context, tx *sqlx.Tx, setting *entity.UserSetting) error {
	sql, args, err := sq.Insert(userSettingTable).
		Columns(userSettingColumns...).
		Values(
			setting.UserID,
			setting.Language,
			setting.CreatedAt,
			setting.UpdatedAt,
		).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, sql, args...); err != nil {
		return err
	}
	return nil
}
