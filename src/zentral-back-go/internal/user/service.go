package user

import "github.com/google/uuid"

// UserService интерфейс для логики пользователя
type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id uuid.UUID) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

// userService структура для логики пользователя
type userService struct {
	repo UserRepository
}

// NewUserService конструктор для создания нового сервиса
func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser создаёт пользователя
func (s *userService) CreateUser(user *User) error {
	return s.repo.Create(user)
}

// GetUserByID возвращает пользователя по ID
func (s *userService) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.FindByID(id.String())
}

// UpdateUser обновляет данные пользователя
func (s *userService) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

// DeleteUser удаляет пользователя
func (s *userService) DeleteUser(user *User) error {
	return s.repo.Delete(user)
}
