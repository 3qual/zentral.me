package collaborator

import (
	"time"

	"github.com/google/uuid"
)

// Collaborator модель для сотрудника (права доступа к папке)
type Collaborator struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	FolderID  uuid.UUID `json:"folder_id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"` // Возможные значения: "owner", "editor", "viewer"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
