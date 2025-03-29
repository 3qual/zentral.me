package collaborator

import "github.com/google/uuid"

// CollaboratorService интерфейс для логики работы с сотрудниками
type CollaboratorService interface {
	CreateCollaborator(collaborator *Collaborator) error
	GetCollaboratorByID(id uuid.UUID) (*Collaborator, error)
	GetCollaboratorsByFolderID(folderID uuid.UUID) ([]Collaborator, error)
	GetCollaboratorsByUserID(userID uuid.UUID) ([]Collaborator, error)
	UpdateCollaborator(collaborator *Collaborator) error
	DeleteCollaborator(collaborator *Collaborator) error
}

// collaboratorService структура для логики работы с сотрудниками
type collaboratorService struct {
	repo CollaboratorRepository
}

// NewCollaboratorService конструктор для создания нового сервиса
func NewCollaboratorService(repo CollaboratorRepository) CollaboratorService {
	return &collaboratorService{
		repo: repo,
	}
}

// CreateCollaborator создает нового сотрудника
func (s *collaboratorService) CreateCollaborator(collaborator *Collaborator) error {
	return s.repo.Create(collaborator)
}

// GetCollaboratorByID возвращает сотрудника по ID
func (s *collaboratorService) GetCollaboratorByID(id uuid.UUID) (*Collaborator, error) {
	return s.repo.FindByID(id.String())
}

// GetCollaboratorsByFolderID возвращает всех сотрудников для папки по её ID
func (s *collaboratorService) GetCollaboratorsByFolderID(folderID uuid.UUID) ([]Collaborator, error) {
	return s.repo.FindByFolderID(folderID.String())
}

// GetCollaboratorsByUserID возвращает всех сотрудников для пользователя по его ID
func (s *collaboratorService) GetCollaboratorsByUserID(userID uuid.UUID) ([]Collaborator, error) {
	return s.repo.FindByUserID(userID.String())
}

// UpdateCollaborator обновляет данные сотрудника
func (s *collaboratorService) UpdateCollaborator(collaborator *Collaborator) error {
	return s.repo.Update(collaborator)
}

// DeleteCollaborator удаляет сотрудника
func (s *collaboratorService) DeleteCollaborator(collaborator *Collaborator) error {
	return s.repo.Delete(collaborator)
}
