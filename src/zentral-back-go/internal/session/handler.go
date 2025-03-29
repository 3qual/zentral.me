package session

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// SessionHandler структура для обработчика сессий
type SessionHandler struct {
	service SessionService
}

// NewSessionHandler конструктор для создания нового обработчика
func NewSessionHandler(service SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

// CreateSessionHandler создаёт новую сессию
func (h *SessionHandler) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	var session Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateSession(&session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(session)
}

// GetSessionByIDHandler получает сессию по ID
func (h *SessionHandler) GetSessionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	session, err := h.service.GetSessionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)
}

// GetSessionsByUserIDHandler получает все сессии по ID пользователя
func (h *SessionHandler) GetSessionsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	sessions, err := h.service.GetSessionsByUserID(uuid.Must(uuid.Parse(userID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessions)
}

// UpdateSessionHandler обновляет данные сессии
func (h *SessionHandler) UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	var session Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateSession(&session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)
}

// DeleteSessionHandler удаляет сессию
func (h *SessionHandler) DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	session, err := h.service.GetSessionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteSession(session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
