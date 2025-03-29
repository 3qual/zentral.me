package refreshtoken

import (
	"gorm.io/gorm"
)

// RefreshTokenRepository интерфейс для работы с данными токенов обновления
type RefreshTokenRepository interface {
	FindByID(id string) (*RefreshToken, error)
	FindByUserID(userID string) ([]RefreshToken, error)
	Create(refreshToken *RefreshToken) error
	Update(refreshToken *RefreshToken) error
	Delete(refreshToken *RefreshToken) error
}

// refreshtokenRepository структура для реализации репозитория
type refreshtokenRepository struct {
	DB *gorm.DB
}

// NewRefreshTokenRepository конструктор для создания нового репозитория
func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshtokenRepository{
		DB: db,
	}
}

// FindByID находит токен обновления по ID
func (r *refreshtokenRepository) FindByID(id string) (*RefreshToken, error) {
	var refreshToken RefreshToken
	err := r.DB.First(&refreshToken, "id = ?", id).Error
	return &refreshToken, err
}

// FindByUserID находит все токены обновления по ID пользователя
func (r *refreshtokenRepository) FindByUserID(userID string) ([]RefreshToken, error) {
	var refreshTokens []RefreshToken
	err := r.DB.Find(&refreshTokens, "user_id = ?", userID).Error
	return refreshTokens, err
}

// Create создает новый токен обновления
func (r *refreshtokenRepository) Create(refreshToken *RefreshToken) error {
	return r.DB.Create(refreshToken).Error
}

// Update обновляет данные токена обновления
func (r *refreshtokenRepository) Update(refreshToken *RefreshToken) error {
	return r.DB.Save(refreshToken).Error
}

// Delete удаляет токен обновления
func (r *refreshtokenRepository) Delete(refreshToken *RefreshToken) error {
	return r.DB.Delete(refreshToken).Error
}
