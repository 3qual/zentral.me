package accesstoken

import "github.com/google/uuid"

// AccessTokenService интерфейс для логики работы с токенами доступа
type AccessTokenService interface {
	CreateAccessToken(accessToken *AccessToken) error
	GetAccessTokenByID(id uuid.UUID) (*AccessToken, error)
	GetAccessTokensByUserID(userID uuid.UUID) ([]AccessToken, error)
	UpdateAccessToken(accessToken *AccessToken) error
	DeleteAccessToken(accessToken *AccessToken) error
}

// accessTokenService структура для логики работы с токенами доступа
type accessTokenService struct {
	repo AccessTokenRepository
}

// NewAccessTokenService конструктор для создания нового сервиса
func NewAccessTokenService(repo AccessTokenRepository) AccessTokenService {
	return &accessTokenService{
		repo: repo,
	}
}

// CreateAccessToken создает токен доступа
func (s *accessTokenService) CreateAccessToken(accessToken *AccessToken) error {
	return s.repo.Create(accessToken)
}

// GetAccessTokenByID возвращает токен доступа по ID
func (s *accessTokenService) GetAccessTokenByID(id uuid.UUID) (*AccessToken, error) {
	return s.repo.FindByID(id.String())
}

// GetAccessTokensByUserID возвращает все токены доступа по ID пользователя
func (s *accessTokenService) GetAccessTokensByUserID(userID uuid.UUID) ([]AccessToken, error) {
	return s.repo.FindByUserID(userID.String())
}

// UpdateAccessToken обновляет данные токена доступа
func (s *accessTokenService) UpdateAccessToken(accessToken *AccessToken) error {
	return s.repo.Update(accessToken)
}

// DeleteAccessToken удаляет токен доступа
func (s *accessTokenService) DeleteAccessToken(accessToken *AccessToken) error {
	return s.repo.Delete(accessToken)
}
