package collaborator

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// CollaboratorHandler структура для обработчика сотрудников
type CollaboratorHandler struct {
	service CollaboratorService
}

// NewCollaboratorHandler конструктор для создания нового обработчика
func NewCollaboratorHandler(service CollaboratorService) *CollaboratorHandler {
	return &CollaboratorHandler{service: service}
}

// CreateCollaboratorHandler создает нового сотрудника
func (h *CollaboratorHandler) CreateCollaboratorHandler(w http.ResponseWriter, r *http.Request) {
	var collaborator Collaborator
	if err := json.NewDecoder(r.Body).Decode(&collaborator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCollaborator(&collaborator); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(collaborator)
}

// GetCollaboratorByIDHandler получает сотрудника по ID
func (h *CollaboratorHandler) GetCollaboratorByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	collaborator, err := h.service.GetCollaboratorByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collaborator)
}

// GetCollaboratorsByFolderIDHandler получает всех сотрудников по ID папки
func (h *CollaboratorHandler) GetCollaboratorsByFolderIDHandler(w http.ResponseWriter, r *http.Request) {
	folderID := chi.URLParam(r, "folder_id")
	collaborators, err := h.service.GetCollaboratorsByFolderID(uuid.Must(uuid.Parse(folderID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collaborators)
}

// GetCollaboratorsByUserIDHandler получает всех сотрудников по ID пользователя
func (h *CollaboratorHandler) GetCollaboratorsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	collaborators, err := h.service.GetCollaboratorsByUserID(uuid.Must(uuid.Parse(userID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collaborators)
}

// UpdateCollaboratorHandler обновляет данные сотрудника
func (h *CollaboratorHandler) UpdateCollaboratorHandler(w http.ResponseWriter, r *http.Request) {
	var collaborator Collaborator
	if err := json.NewDecoder(r.Body).Decode(&collaborator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCollaborator(&collaborator); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collaborator)
}

// DeleteCollaboratorHandler удаляет сотрудника
func (h *CollaboratorHandler) DeleteCollaboratorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	collaborator, err := h.service.GetCollaboratorByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteCollaborator(collaborator); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
