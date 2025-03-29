package user

import (
	"time"

	"github.com/google/uuid"
)

// User модель для пользователя
type User struct {
	ID              uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	IsActive        bool      `json:"is_active"`
	IsVisible       bool      `json:"is_visible"`
	Name            string    `json:"name"`
	Username        string    `gorm:"unique" json:"username"`
	PasswordHash    string    `json:"password_hash"`
	Email           string    `gorm:"unique" json:"email"`
	IsVerified      bool      `json:"is_verified"`
	PrimaryCurrency string    `json:"primary_currency"`
	AvatarPath      string    `json:"avatar_path"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
