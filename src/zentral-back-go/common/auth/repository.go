package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	u "github.com/3qual/zentral-back-go/internal/user"
)

// AuthRepository интерфейс для работы с данными аутентификации
type AuthRepository interface {
	FindByEmail(email string) (*u.User, error)
	FindByID(id uuid.UUID) (*u.User, error)
	CreateUser(user *u.User) error
}

// authRepository структура для реализации репозитория
type authRepository struct {
	DB *gorm.DB
}

// NewAuthRepository конструктор для создания нового репозитория
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		DB: db,
	}
}

// FindByEmail находит пользователя по email
func (r *authRepository) FindByEmail(email string) (*u.User, error) {
	var user u.User
	err := r.DB.First(&user, "email = ?", email).Error
	return &user, err
}

// FindByID находит пользователя по ID
func (r *authRepository) FindByID(id uuid.UUID) (*u.User, error) {
	var user u.User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err
}

// CreateUser создает нового пользователя
func (r *authRepository) CreateUser(user *u.User) error {
	return r.DB.Create(user).Error
}
