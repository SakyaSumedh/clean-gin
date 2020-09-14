package models

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	Email        *string        `json:"email"`
	Age          uint8          `json:"age"`
	Birthday     *time.Time     `json:"time"`
	MemberNumber sql.NullString `json:"member_number"`
	ActivedAt    sql.NullTime   `json:"active_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
