package conv

import (
	"database/sql"
	"time"
)

func NullTimeToTime(nullTime sql.NullTime) *time.Time {
	if !nullTime.Valid {
		return nil
	}
	v := nullTime.Time
	return &v
}
