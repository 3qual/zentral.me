package foldertransaction

import "github.com/google/uuid"

// FolderTransactionService интерфейс для логики работы с папками и транзакциями
type FolderTransactionService interface {
	CreateFolderTransaction(folderTransaction *FolderTransaction) error
	GetFolderTransactionByID(id uuid.UUID) (*FolderTransaction, error)
	GetFolderTransactionsByFolderID(folderID uuid.UUID) ([]FolderTransaction, error)
	GetFolderTransactionsByTransactionID(transactionID uuid.UUID) ([]FolderTransaction, error)
	DeleteFolderTransaction(folderTransaction *FolderTransaction) error
}

// folderTransactionService структура для логики работы с папками и транзакциями
type folderTransactionService struct {
	repo FolderTransactionRepository
}

// NewFolderTransactionService конструктор для создания нового сервиса
func NewFolderTransactionService(repo FolderTransactionRepository) FolderTransactionService {
	return &folderTransactionService{
		repo: repo,
	}
}

// CreateFolderTransaction создает связь между папкой и транзакцией
func (s *folderTransactionService) CreateFolderTransaction(folderTransaction *FolderTransaction) error {
	return s.repo.Create(folderTransaction)
}

// GetFolderTransactionByID возвращает связь по ID
func (s *folderTransactionService) GetFolderTransactionByID(id uuid.UUID) (*FolderTransaction, error) {
	return s.repo.FindByID(id.String())
}

// GetFolderTransactionsByFolderID возвращает все связи для папки по её ID
func (s *folderTransactionService) GetFolderTransactionsByFolderID(folderID uuid.UUID) ([]FolderTransaction, error) {
	return s.repo.FindByFolderID(folderID.String())
}

// GetFolderTransactionsByTransactionID возвращает все связи для транзакции по её ID
func (s *folderTransactionService) GetFolderTransactionsByTransactionID(transactionID uuid.UUID) ([]FolderTransaction, error) {
	return s.repo.FindByTransactionID(transactionID.String())
}

// DeleteFolderTransaction удаляет связь
func (s *folderTransactionService) DeleteFolderTransaction(folderTransaction *FolderTransaction) error {
	return s.repo.Delete(folderTransaction)
}
