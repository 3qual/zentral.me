package folder

import (
	"time"

	"github.com/google/uuid"
)

// Folder модель для папки
type Folder struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`      // Неуникальное название папки
	IconPath  string    `json:"icon_path"` // Путь к файлу иконки папки
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
