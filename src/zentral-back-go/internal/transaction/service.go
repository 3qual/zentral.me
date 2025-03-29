package transaction

import "github.com/google/uuid"

// TransactionService интерфейс для логики транзакции
type TransactionService interface {
	CreateTransaction(transaction *Transaction) error
	GetTransactionByID(id uuid.UUID) (*Transaction, error)
	GetTransactionByUserID(userID uuid.UUID, page int, limit int) ([]Transaction, error)
	GetAllTransactions() ([]Transaction, error)
	UpdateTransaction(transaction *Transaction) error
	DeleteTransaction(transaction *Transaction) error
}

// transactionService структура для логики транзакции
type transactionService struct {
	repo TransactionRepository
}

// NewTransactionService конструктор для создания нового сервиса
func NewTransactionService(repo TransactionRepository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

// CreateTransaction создает транзакцию
func (s *transactionService) CreateTransaction(transaction *Transaction) error {
	return s.repo.Create(transaction)
}

// GetTransactionByID возвращает транзакцию по ID
func (s *transactionService) GetTransactionByID(id uuid.UUID) (*Transaction, error) {
	return s.repo.FindByID(id.String())
}

// GetTransactionByUserID возвращает все транзакции по ID пользователя с пагинацией
func (s *transactionService) GetTransactionByUserID(userID uuid.UUID, page int, limit int) ([]Transaction, error) {
	return s.repo.FindByUserID(userID.String(), page, limit)
}

// GetAllTransactions возвращает все транзакции
func (s *transactionService) GetAllTransactions() ([]Transaction, error) {
	return s.repo.FindAll()
}

// UpdateTransaction обновляет данные транзакции
func (s *transactionService) UpdateTransaction(transaction *Transaction) error {
	return s.repo.Update(transaction)
}

// DeleteTransaction удаляет транзакцию
func (s *transactionService) DeleteTransaction(transaction *Transaction) error {
	return s.repo.Delete(transaction)
}
