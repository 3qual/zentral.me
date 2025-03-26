package accesstoken

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// AccessTokenHandler структура для обработчика токенов доступа
type AccessTokenHandler struct {
	service AccessTokenService
}

// NewAccessTokenHandler конструктор для создания нового обработчика
func NewAccessTokenHandler(service AccessTokenService) *AccessTokenHandler {
	return &AccessTokenHandler{service: service}
}

// CreateAccessTokenHandler создаёт новый токен доступа
func (h *AccessTokenHandler) CreateAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var accessToken AccessToken
	if err := json.NewDecoder(r.Body).Decode(&accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateAccessToken(&accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accessToken)
}

// GetAccessTokenByIDHandler получает токен доступа по ID
func (h *AccessTokenHandler) GetAccessTokenByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	accessToken, err := h.service.GetAccessTokenByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// GetAccessTokensByUserIDHandler получает все токены по ID пользователя
func (h *AccessTokenHandler) GetAccessTokensByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	accessTokens, err := h.service.GetAccessTokensByUserID(uuid.Must(uuid.Parse(userID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessTokens)
}

// UpdateAccessTokenHandler обновляет данные токена доступа
func (h *AccessTokenHandler) UpdateAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var accessToken AccessToken
	if err := json.NewDecoder(r.Body).Decode(&accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateAccessToken(&accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// DeleteAccessTokenHandler удаляет токен доступа
func (h *AccessTokenHandler) DeleteAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	accessToken, err := h.service.GetAccessTokenByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteAccessToken(accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
