package model

import (
	"time"
)

// User is a user
type User struct {
	ID    int64  `json:"id" meddler:"user_id,pk"`
	Login string `json:"login" meddler:"user_login"`

	Token  string    `json:"-" meddler:"user_token"`
	Secret string    `json:"-" meddler:"user_secret"`
	Expiry time.Time `json:"-" meddler:"user_expiry,localtime" binding:"required"`

	Email string `json:"email" meddler:"user_email"`

	Hash string `json:"-" meddler:"user_hash"`
}
