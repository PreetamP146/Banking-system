package models

import (
	"time"

	"github.com/google/uuid"
)

// User represent schema for users table in database
type User struct {
	ID            uuid.UUID `json:"id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Email         string    `json:"email"`
	Password_Hash string    `json:"-"`
	Role          string    `json:"role"`
	Is_active     bool      `json:"is_active"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
type RegisterUserRequest struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"-"`
}
