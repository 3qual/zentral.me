package foldertransaction

import (
	"time"

	"github.com/google/uuid"
)

// FolderTransaction модель для связи папок и транзакций
type FolderTransaction struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	FolderID      uuid.UUID `json:"folder_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
