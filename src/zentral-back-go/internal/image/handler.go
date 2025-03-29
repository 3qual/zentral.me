package image

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// ImageHandler структура для обработчика изображений
type ImageHandler struct {
	service ImageService
}

// NewImageHandler конструктор для создания нового обработчика
func NewImageHandler(service ImageService) *ImageHandler {
	return &ImageHandler{service: service}
}

// CreateImageHandler создает новое изображение
func (h *ImageHandler) CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	var image Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateImage(&image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
}

// GetImageByIDHandler получает изображение по ID
func (h *ImageHandler) GetImageByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	image, err := h.service.GetImageByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(image)
}

// GetImagesByTransactionIDHandler получает все изображения по ID транзакции
func (h *ImageHandler) GetImagesByTransactionIDHandler(w http.ResponseWriter, r *http.Request) {
	transactionID := chi.URLParam(r, "transaction_id")
	images, err := h.service.GetImagesByTransactionID(uuid.Must(uuid.Parse(transactionID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}

// DeleteImageHandler удаляет изображение
func (h *ImageHandler) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	image, err := h.service.GetImageByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteImage(image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
