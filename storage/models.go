// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package storage

import (
	"time"
)

type Tag struct {
	ID        int64
	TagName   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
