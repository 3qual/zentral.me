package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	u "github.com/3qual/zentral-back-go/internal/user"
)

// AuthHandler структура для обработчика аутентификации
type AuthHandler struct {
	service AuthService
}

// NewAuthHandler конструктор для создания нового обработчика
func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// RegisterHandler регистрирует нового пользователя
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user u.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Регистрация пользователя
	token, err := h.service.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// LoginHandler выполняет вход пользователя
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials map[string]string
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := credentials["email"]
	password := credentials["password"]

	// Вход в систему
	token, err := h.service.Login(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// GetUserByIDHandler получает пользователя по ID
func (h *AuthHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := uuid.Must(uuid.Parse(id))

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
