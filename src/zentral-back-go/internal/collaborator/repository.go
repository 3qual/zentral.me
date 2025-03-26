package collaborator

import (
	"gorm.io/gorm"
)

// CollaboratorRepository интерфейс для работы с данными сотрудника
type CollaboratorRepository interface {
	FindByID(id string) (*Collaborator, error)
	FindByFolderID(folderID string) ([]Collaborator, error)
	FindByUserID(userID string) ([]Collaborator, error)
	Create(collaborator *Collaborator) error
	Update(collaborator *Collaborator) error
	Delete(collaborator *Collaborator) error
}

// collaboratorRepository структура для реализации репозитория
type collaboratorRepository struct {
	DB *gorm.DB
}

// NewCollaboratorRepository конструктор для создания нового репозитория
func NewCollaboratorRepository(db *gorm.DB) CollaboratorRepository {
	return &collaboratorRepository{
		DB: db,
	}
}

// FindByID находит сотрудника по ID
func (r *collaboratorRepository) FindByID(id string) (*Collaborator, error) {
	var collaborator Collaborator
	err := r.DB.First(&collaborator, "id = ?", id).Error
	return &collaborator, err
}

// FindByFolderID находит всех сотрудников по ID папки
func (r *collaboratorRepository) FindByFolderID(folderID string) ([]Collaborator, error) {
	var collaborators []Collaborator
	err := r.DB.Find(&collaborators, "folder_id = ?", folderID).Error
	return collaborators, err
}

// FindByUserID находит всех сотрудников по ID пользователя
func (r *collaboratorRepository) FindByUserID(userID string) ([]Collaborator, error) {
	var collaborators []Collaborator
	err := r.DB.Find(&collaborators, "user_id = ?", userID).Error
	return collaborators, err
}

// Create создает нового сотрудника
func (r *collaboratorRepository) Create(collaborator *Collaborator) error {
	return r.DB.Create(collaborator).Error
}

// Update обновляет данные сотрудника
func (r *collaboratorRepository) Update(collaborator *Collaborator) error {
	return r.DB.Save(collaborator).Error
}

// Delete удаляет сотрудника
func (r *collaboratorRepository) Delete(collaborator *Collaborator) error {
	return r.DB.Delete(collaborator).Error
}
