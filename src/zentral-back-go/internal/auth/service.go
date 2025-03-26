package auth

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	u "github.com/3qual/zentral-back-go/internal/user"
)

// AuthService интерфейс для логики аутентификации
type AuthService interface {
	Register(user *u.User) (string, error)
	Login(email, password string) (string, error)
	GetUserByID(userID uuid.UUID) (*u.User, error)
}

// authService структура для логики аутентификации
type authService struct {
	repo AuthRepository
}

// NewAuthService конструктор для создания нового сервиса
func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

// Register регистрирует нового пользователя
func (s *authService) Register(user *u.User) (string, error) {
	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(hashedPassword)

	// Создание пользователя в БД
	if err := s.repo.CreateUser(user); err != nil {
		return "", err
	}

	// Генерация JWT токена
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login выполняет вход пользователя
func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Генерация JWT токена
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID возвращает пользователя по ID
func (s *authService) GetUserByID(userID uuid.UUID) (*u.User, error) {
	return s.repo.FindByID(userID)
}
