package transaction

import (
	"gorm.io/gorm"
)

// TransactionRepository интерфейс для работы с данными транзакции
type TransactionRepository interface {
	FindByID(id string) (*Transaction, error)
	FindByUserID(userID string, page int, limit int) ([]Transaction, error)
	FindAll() ([]Transaction, error)
	Create(transaction *Transaction) error
	Update(transaction *Transaction) error
	Delete(transaction *Transaction) error
}

// transactionRepository структура для реализации репозитория
type transactionRepository struct {
	DB *gorm.DB
}

// NewTransactionRepository конструктор для создания нового репозитория
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

// FindByID находит транзакцию по ID
func (r *transactionRepository) FindByID(id string) (*Transaction, error) {
	var transaction Transaction
	err := r.DB.First(&transaction, "id = ?", id).Error
	return &transaction, err
}

// FindByUserID возвращает все транзакции по ID пользователя с пагинацией
func (r *transactionRepository) FindByUserID(userID string, page int, limit int) ([]Transaction, error) {
	var transactions []Transaction

	// Используем пагинацию, начиная с (page-1)*limit
	offset := (page - 1) * limit

	err := r.DB.Limit(limit).Offset(offset).Find(&transactions, "user_id = ?", userID).Error
	return transactions, err
}

// FindAll находит все транзакции
func (r *transactionRepository) FindAll() ([]Transaction, error) {
	var transactions []Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}

// Create создает новую транзакцию
func (r *transactionRepository) Create(transaction *Transaction) error {
	return r.DB.Create(transaction).Error
}

// Update обновляет данные транзакции
func (r *transactionRepository) Update(transaction *Transaction) error {
	return r.DB.Save(transaction).Error
}

// Delete удаляет транзакцию
func (r *transactionRepository) Delete(transaction *Transaction) error {
	return r.DB.Delete(transaction).Error
}
