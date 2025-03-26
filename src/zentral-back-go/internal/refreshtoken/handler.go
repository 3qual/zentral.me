package refreshtoken

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// RefreshTokenHandler структура для обработчика токенов обновления
type RefreshTokenHandler struct {
	service RefreshTokenService
}

// NewRefreshTokenHandler конструктор для создания нового обработчика
func NewRefreshTokenHandler(service RefreshTokenService) *RefreshTokenHandler {
	return &RefreshTokenHandler{service: service}
}

// CreateRefreshTokenHandler создает новый токен обновления
func (h *RefreshTokenHandler) CreateRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var refreshToken RefreshToken
	if err := json.NewDecoder(r.Body).Decode(&refreshToken); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateRefreshToken(&refreshToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(refreshToken)
}

// GetRefreshTokenByIDHandler получает токен обновления по ID
func (h *RefreshTokenHandler) GetRefreshTokenByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	refreshToken, err := h.service.GetRefreshTokenByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refreshToken)
}

// GetRefreshTokensByUserIDHandler получает все токены обновления по ID пользователя
func (h *RefreshTokenHandler) GetRefreshTokensByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	refreshTokens, err := h.service.GetRefreshTokensByUserID(uuid.Must(uuid.Parse(userID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refreshTokens)
}

// UpdateRefreshTokenHandler обновляет данные токена обновления
func (h *RefreshTokenHandler) UpdateRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var refreshToken RefreshToken
	if err := json.NewDecoder(r.Body).Decode(&refreshToken); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateRefreshToken(&refreshToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refreshToken)
}

// DeleteRefreshTokenHandler удаляет токен обновления
func (h *RefreshTokenHandler) DeleteRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	refreshToken, err := h.service.GetRefreshTokenByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteRefreshToken(refreshToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
