package refreshtoken

import "github.com/google/uuid"

// RefreshTokenService интерфейс для логики работы с токенами обновления
type RefreshTokenService interface {
	CreateRefreshToken(refreshToken *RefreshToken) error
	GetRefreshTokenByID(id uuid.UUID) (*RefreshToken, error)
	GetRefreshTokensByUserID(userID uuid.UUID) ([]RefreshToken, error)
	UpdateRefreshToken(refreshToken *RefreshToken) error
	DeleteRefreshToken(refreshToken *RefreshToken) error
}

// refreshtokenService структура для логики работы с токенами обновления
type refreshtokenService struct {
	repo RefreshTokenRepository
}

// NewRefreshTokenService конструктор для создания нового сервиса
func NewRefreshTokenService(repo RefreshTokenRepository) RefreshTokenService {
	return &refreshtokenService{
		repo: repo,
	}
}

// CreateRefreshToken создает токен обновления
func (s *refreshtokenService) CreateRefreshToken(refreshToken *RefreshToken) error {
	return s.repo.Create(refreshToken)
}

// GetRefreshTokenByID возвращает токен обновления по ID
func (s *refreshtokenService) GetRefreshTokenByID(id uuid.UUID) (*RefreshToken, error) {
	return s.repo.FindByID(id.String())
}

// GetRefreshTokensByUserID возвращает все токены обновления по ID пользователя
func (s *refreshtokenService) GetRefreshTokensByUserID(userID uuid.UUID) ([]RefreshToken, error) {
	return s.repo.FindByUserID(userID.String())
}

// UpdateRefreshToken обновляет данные токена обновления
func (s *refreshtokenService) UpdateRefreshToken(refreshToken *RefreshToken) error {
	return s.repo.Update(refreshToken)
}

// DeleteRefreshToken удаляет токен обновления
func (s *refreshtokenService) DeleteRefreshToken(refreshToken *RefreshToken) error {
	return s.repo.Delete(refreshToken)
}
