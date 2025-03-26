package transaction

import (
	"time"

	"github.com/google/uuid"
)

// Transaction модель для транзакции
type Transaction struct {
	ID                  uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Type                string    `json:"type"`
	Name                string    `json:"name"`
	Date                time.Time `json:"date"`
	MerchantName        string    `json:"merchant_name"`
	TotalAmount         float64   `json:"total_amount"`
	Currency            string    `json:"currency"`
	RecognizedText      string    `json:"recognized_text"`
	Description         string    `json:"description"`
	Category            string    `json:"category"`
	LocationName        string    `json:"location_name"`
	LocationAddress     string    `json:"location_address"`
	LocationCoordinates string    `json:"location_coordinates"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
