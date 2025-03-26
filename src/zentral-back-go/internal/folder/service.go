package folder

import "github.com/google/uuid"

// FolderService интерфейс для логики папки
type FolderService interface {
	CreateFolder(folder *Folder) error
	GetFolderByID(id uuid.UUID) (*Folder, error)
	GetAllFolders() ([]Folder, error)
	UpdateFolder(folder *Folder) error
	DeleteFolder(folder *Folder) error
}

// folderService структура для логики папки
type folderService struct {
	repo FolderRepository
}

// NewFolderService конструктор для создания нового сервиса
func NewFolderService(repo FolderRepository) FolderService {
	return &folderService{
		repo: repo,
	}
}

// CreateFolder создает папку
func (s *folderService) CreateFolder(folder *Folder) error {
	return s.repo.Create(folder)
}

// GetFolderByID возвращает папку по ID
func (s *folderService) GetFolderByID(id uuid.UUID) (*Folder, error) {
	return s.repo.FindByID(id.String())
}

// GetAllFolders возвращает все папки
func (s *folderService) GetAllFolders() ([]Folder, error) {
	return s.repo.FindAll()
}

// UpdateFolder обновляет данные папки
func (s *folderService) UpdateFolder(folder *Folder) error {
	return s.repo.Update(folder)
}

// DeleteFolder удаляет папку
func (s *folderService) DeleteFolder(folder *Folder) error {
	return s.repo.Delete(folder)
}
