package entity

import (
	"database/sql"
	"time"
)

type Event struct {
	ID               int64        `db:"id"`
	OwnerID          int64        `db:"owner_id"`
	Title            string       `db:"title"`
	Description      string       `db:"description"`
	TimeStart        time.Time    `db:"time_start"`
	TimeEnd          time.Time    `db:"time_end"`
	NotificationTime sql.NullTime `db:"notification_time"`
}
