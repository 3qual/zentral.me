package session

import "github.com/google/uuid"

// SessionService интерфейс для логики работы с сессиями
type SessionService interface {
	CreateSession(session *Session) error
	GetSessionByID(id uuid.UUID) (*Session, error)
	GetSessionsByUserID(userID uuid.UUID) ([]Session, error)
	UpdateSession(session *Session) error
	DeleteSession(session *Session) error
}

// sessionService структура для логики работы с сессиями
type sessionService struct {
	repo SessionRepository
}

// NewSessionService конструктор для создания нового сервиса
func NewSessionService(repo SessionRepository) SessionService {
	return &sessionService{
		repo: repo,
	}
}

// CreateSession создает сессию
func (s *sessionService) CreateSession(session *Session) error {
	return s.repo.Create(session)
}

// GetSessionByID возвращает сессию по ID
func (s *sessionService) GetSessionByID(id uuid.UUID) (*Session, error) {
	return s.repo.FindByID(id.String())
}

// GetSessionsByUserID возвращает все сессии по ID пользователя
func (s *sessionService) GetSessionsByUserID(userID uuid.UUID) ([]Session, error) {
	return s.repo.FindByUserID(userID.String())
}

// UpdateSession обновляет данные сессии
func (s *sessionService) UpdateSession(session *Session) error {
	return s.repo.Update(session)
}

// DeleteSession удаляет сессию
func (s *sessionService) DeleteSession(session *Session) error {
	return s.repo.Delete(session)
}
