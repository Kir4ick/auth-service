package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id    int       `db:"id"`
	Uuid  uuid.UUID `db:"uuid"`
	Login string    `db:"login"`
	Email string    `db:"email"`
	Hash  string    `db:"hash"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	EmailVerifiedAt *time.Time `db:"email_verified_at"`
}
