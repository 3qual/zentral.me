package session

import (
	"gorm.io/gorm"
)

// SessionRepository интерфейс для работы с данными сессий
type SessionRepository interface {
	FindByID(id string) (*Session, error)
	FindByUserID(userID string) ([]Session, error)
	Create(session *Session) error
	Update(session *Session) error
	Delete(session *Session) error
}

// sessionRepository структура для реализации репозитория
type sessionRepository struct {
	DB *gorm.DB
}

// NewSessionRepository конструктор для создания нового репозитория
func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{
		DB: db,
	}
}

// FindByID находит сессию по ID
func (r *sessionRepository) FindByID(id string) (*Session, error) {
	var session Session
	err := r.DB.First(&session, "id = ?", id).Error
	return &session, err
}

// FindByUserID находит все сессии по ID пользователя
func (r *sessionRepository) FindByUserID(userID string) ([]Session, error) {
	var sessions []Session
	err := r.DB.Find(&sessions, "user_id = ?", userID).Error
	return sessions, err
}

// Create создает новую сессию
func (r *sessionRepository) Create(session *Session) error {
	return r.DB.Create(session).Error
}

// Update обновляет данные сессии
func (r *sessionRepository) Update(session *Session) error {
	return r.DB.Save(session).Error
}

// Delete удаляет сессию
func (r *sessionRepository) Delete(session *Session) error {
	return r.DB.Delete(session).Error
}
