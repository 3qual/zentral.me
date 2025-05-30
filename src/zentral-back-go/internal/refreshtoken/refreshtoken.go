package refreshtoken

import (
	"time"

	"github.com/google/uuid"
)

// RefreshToken модель для токена обновления
type RefreshToken struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Token      string    `json:"token"`
	ExpiryDate time.Time `json:"expiry_date"`
	SessionID  uuid.UUID `json:"session_id"`
	Revoked    bool      `json:"revoked"` // Если true, токен недействителен
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
