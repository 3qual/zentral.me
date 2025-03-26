package folder

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// FolderHandler структура для обработчика папки
type FolderHandler struct {
	service FolderService
}

// NewFolderHandler конструктор для создания нового обработчика
func NewFolderHandler(service FolderService) *FolderHandler {
	return &FolderHandler{service: service}
}

// CreateFolderHandler создает новую папку
func (h *FolderHandler) CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
	var folder Folder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateFolder(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(folder)
}

// GetFolderByIDHandler получает папку по ID
func (h *FolderHandler) GetFolderByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	folder, err := h.service.GetFolderByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folder)
}

// UpdateFolderHandler обновляет данные папки
func (h *FolderHandler) UpdateFolderHandler(w http.ResponseWriter, r *http.Request) {
	var folder Folder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateFolder(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folder)
}

// DeleteFolderHandler удаляет папку
func (h *FolderHandler) DeleteFolderHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	f, err := h.service.GetFolderByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteFolder(f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
