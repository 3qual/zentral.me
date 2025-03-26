package folder

import (
	"gorm.io/gorm"
)

// FolderRepository интерфейс для работы с данными папки
type FolderRepository interface {
	FindByID(id string) (*Folder, error)
	FindAll() ([]Folder, error)
	Create(folder *Folder) error
	Update(folder *Folder) error
	Delete(folder *Folder) error
}

// folderRepository структура для реализации репозитория
type folderRepository struct {
	DB *gorm.DB
}

// NewFolderRepository конструктор для создания нового репозитория
func NewFolderRepository(db *gorm.DB) FolderRepository {
	return &folderRepository{
		DB: db,
	}
}

// FindByID находит папку по ID
func (r *folderRepository) FindByID(id string) (*Folder, error) {
	var folder Folder
	err := r.DB.First(&folder, "id = ?", id).Error
	return &folder, err
}

// FindAll находит все папки
func (r *folderRepository) FindAll() ([]Folder, error) {
	var folders []Folder
	err := r.DB.Find(&folders).Error
	return folders, err
}

// Create создает новую папку
func (r *folderRepository) Create(folder *Folder) error {
	return r.DB.Create(folder).Error
}

// Update обновляет данные папки
func (r *folderRepository) Update(folder *Folder) error {
	return r.DB.Save(folder).Error
}

// Delete удаляет папку
func (r *folderRepository) Delete(folder *Folder) error {
	return r.DB.Delete(folder).Error
}
