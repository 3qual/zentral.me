package session

import (
	"time"

	"github.com/google/uuid"
)

// Session модель для сессии пользователя
type Session struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	DeviceName   string    `json:"device_name"`   // Имя устройства
	DeviceType   string    `json:"device_type"`   // Тип устройства (например, Web, iPhone, Android)
	DeviceIP     string    `json:"device_ip"`     // IP адрес устройства
	LastActivity time.Time `json:"last_activity"` // Время последней активности
	Active       bool      `json:"active"`        // Статус активности сессии
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
