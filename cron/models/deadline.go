package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Deadline struct {
	ID            uuid.UUID `db:"id"`
	RemindTo      string    `db:"remind_to"`
	DaysRemaining int32     `db:"days_remaining"`
	Warned        bool      `db:"warned"`
	Deadline      time.Time `db:"deadline"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
