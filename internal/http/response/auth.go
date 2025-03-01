package response

import (
	"github.com/google/uuid"
	"time"
)

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Id              int        `json:"id"`
	Uuid            uuid.UUID  `json:"uuid"`
	Login           string     `json:"login"`
	Email           string     `json:"email"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
}

type SignUpResponse struct {
	User
}
