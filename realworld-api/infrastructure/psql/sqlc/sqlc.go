package sqlc

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func toTimestamptz(t time.Time) pgtype.Timestamptz {
	if t == (time.Time{}) {
		return pgtype.Timestamptz{
			Valid: false,
		}
	}
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
