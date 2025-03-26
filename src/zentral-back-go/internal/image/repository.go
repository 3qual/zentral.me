package image

import (
	"gorm.io/gorm"
)

// ImageRepository интерфейс для работы с данными изображений
type ImageRepository interface {
	FindByID(id string) (*Image, error)
	FindByTransactionID(transactionID string) ([]Image, error)
	Create(image *Image) error
	Delete(image *Image) error
}

// imageRepository структура для реализации репозитория
type imageRepository struct {
	DB *gorm.DB
}

// NewImageRepository конструктор для создания нового репозитория
func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{
		DB: db,
	}
}

// FindByID находит изображение по ID
func (r *imageRepository) FindByID(id string) (*Image, error) {
	var image Image
	err := r.DB.First(&image, "id = ?", id).Error
	return &image, err
}

// FindByTransactionID находит все изображения по ID транзакции
func (r *imageRepository) FindByTransactionID(transactionID string) ([]Image, error) {
	var images []Image
	err := r.DB.Find(&images, "transaction_id = ?", transactionID).Error
	return images, err
}

// Create создает новое изображение
func (r *imageRepository) Create(image *Image) error {
	return r.DB.Create(image).Error
}

// Delete удаляет изображение
func (r *imageRepository) Delete(image *Image) error {
	return r.DB.Delete(image).Error
}
