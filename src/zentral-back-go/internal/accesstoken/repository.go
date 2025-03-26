package accesstoken

import (
	"gorm.io/gorm"
)

// AccessTokenRepository интерфейс для работы с данными токенов доступа
type AccessTokenRepository interface {
	FindByID(id string) (*AccessToken, error)
	FindByUserID(userID string) ([]AccessToken, error)
	Create(accessToken *AccessToken) error
	Update(accessToken *AccessToken) error
	Delete(accessToken *AccessToken) error
}

// accessTokenRepository структура для реализации репозитория
type accessTokenRepository struct {
	DB *gorm.DB
}

// NewAccessTokenRepository конструктор для создания нового репозитория
func NewAccessTokenRepository(db *gorm.DB) AccessTokenRepository {
	return &accessTokenRepository{
		DB: db,
	}
}

// FindByID находит токен по ID
func (r *accessTokenRepository) FindByID(id string) (*AccessToken, error) {
	var accessToken AccessToken
	err := r.DB.First(&accessToken, "id = ?", id).Error
	return &accessToken, err
}

// FindByUserID находит все токены по ID пользователя
func (r *accessTokenRepository) FindByUserID(userID string) ([]AccessToken, error) {
	var accessTokens []AccessToken
	err := r.DB.Find(&accessTokens, "user_id = ?", userID).Error
	return accessTokens, err
}

// Create создает новый токен
func (r *accessTokenRepository) Create(accessToken *AccessToken) error {
	return r.DB.Create(accessToken).Error
}

// Update обновляет данные токена
func (r *accessTokenRepository) Update(accessToken *AccessToken) error {
	return r.DB.Save(accessToken).Error
}

// Delete удаляет токен
func (r *accessTokenRepository) Delete(accessToken *AccessToken) error {
	return r.DB.Delete(accessToken).Error
}
