package user

import (
	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с данными пользователя
type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}

// userRepository структура для реализации репозитория
type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository конструктор для создания нового репозитория
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// FindByID находит пользователя по ID
func (r *userRepository) FindByID(id string) (*User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err
}

// FindByEmail находит пользователя по email
func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.DB.First(&user, "email = ?", email).Error
	return &user, err
}

// FindByUsername находит пользователя по username
func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.DB.First(&user, "username = ?", username).Error
	return &user, err
}

// Create создает нового пользователя
func (r *userRepository) Create(user *User) error {
	return r.DB.Create(user).Error
}

// Update обновляет данные пользователя
func (r *userRepository) Update(user *User) error {
	return r.DB.Save(user).Error
}

// Delete удаляет пользователя
func (r *userRepository) Delete(user *User) error {
	return r.DB.Delete(user).Error
}
