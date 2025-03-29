package foldertransaction

import (
	"gorm.io/gorm"
)

// FolderTransactionRepository интерфейс для работы с данными связи папки и транзакции
type FolderTransactionRepository interface {
	FindByID(id string) (*FolderTransaction, error)
	FindByFolderID(folderID string) ([]FolderTransaction, error)
	FindByTransactionID(transactionID string) ([]FolderTransaction, error)
	Create(folderTransaction *FolderTransaction) error
	Delete(folderTransaction *FolderTransaction) error
}

// folderTransactionRepository структура для реализации репозитория
type folderTransactionRepository struct {
	DB *gorm.DB
}

// NewFolderTransactionRepository конструктор для создания нового репозитория
func NewFolderTransactionRepository(db *gorm.DB) FolderTransactionRepository {
	return &folderTransactionRepository{
		DB: db,
	}
}

// FindByID находит связь по ID
func (r *folderTransactionRepository) FindByID(id string) (*FolderTransaction, error) {
	var folderTransaction FolderTransaction
	err := r.DB.First(&folderTransaction, "id = ?", id).Error
	return &folderTransaction, err
}

// FindByFolderID находит все связи по ID папки
func (r *folderTransactionRepository) FindByFolderID(folderID string) ([]FolderTransaction, error) {
	var folderTransactions []FolderTransaction
	err := r.DB.Find(&folderTransactions, "folder_id = ?", folderID).Error
	return folderTransactions, err
}

// FindByTransactionID находит все связи по ID транзакции
func (r *folderTransactionRepository) FindByTransactionID(transactionID string) ([]FolderTransaction, error) {
	var folderTransactions []FolderTransaction
	err := r.DB.Find(&folderTransactions, "transaction_id = ?", transactionID).Error
	return folderTransactions, err
}

// Create создает новую связь
func (r *folderTransactionRepository) Create(folderTransaction *FolderTransaction) error {
	return r.DB.Create(folderTransaction).Error
}

// Delete удаляет связь
func (r *folderTransactionRepository) Delete(folderTransaction *FolderTransaction) error {
	return r.DB.Delete(folderTransaction).Error
}
