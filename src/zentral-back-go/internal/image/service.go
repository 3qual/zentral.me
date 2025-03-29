package image

import "github.com/google/uuid"

// ImageService интерфейс для логики работы с изображениями
type ImageService interface {
	CreateImage(image *Image) error
	GetImageByID(id uuid.UUID) (*Image, error)
	GetImagesByTransactionID(transactionID uuid.UUID) ([]Image, error)
	DeleteImage(image *Image) error
}

// imageService структура для логики работы с изображениями
type imageService struct {
	repo ImageRepository
}

// NewImageService конструктор для создания нового сервиса
func NewImageService(repo ImageRepository) ImageService {
	return &imageService{
		repo: repo,
	}
}

// CreateImage создает изображение
func (s *imageService) CreateImage(image *Image) error {
	return s.repo.Create(image)
}

// GetImageByID возвращает изображение по ID
func (s *imageService) GetImageByID(id uuid.UUID) (*Image, error) {
	return s.repo.FindByID(id.String())
}

// GetImagesByTransactionID возвращает все изображения по ID транзакции
func (s *imageService) GetImagesByTransactionID(transactionID uuid.UUID) ([]Image, error) {
	return s.repo.FindByTransactionID(transactionID.String())
}

// DeleteImage удаляет изображение
func (s *imageService) DeleteImage(image *Image) error {
	return s.repo.Delete(image)
}
