package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID  `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	Email        string     `db:"email" json:"email" validate:"required,email,lte=255"`
	Name         string     `db:"name" json:"name" validate:"required"`
	PasswordHash string     `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserRole     string     `db:"user_role" json:"user_role" validate:"required,lte=25"`
}
