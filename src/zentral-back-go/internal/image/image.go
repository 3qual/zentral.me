package image

import (
	"time"

	"github.com/google/uuid"
)

// Image модель для изображения
type Image struct {
	ID                 uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	TransactionID      uuid.UUID `json:"transaction_id"`
	OriginalImagePath  string    `json:"original_image_path"`  // Путь к оригинальному изображению
	ProcessedImagePath string    `json:"processed_image_path"` // Путь к обработанному изображению
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
