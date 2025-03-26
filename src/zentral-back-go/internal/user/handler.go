package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// UserHandler структура для обработчика пользователя
type UserHandler struct {
	service UserService
}

// NewUserHandler конструктор для создания нового обработчика
func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUserHandler создаёт нового пользователя
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByIDHandler получает пользователя по ID
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.service.GetUserByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler обновляет данные пользователя
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler удаляет пользователя
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.service.GetUserByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UserRouter настраивает маршруты для пользователя
func UserRouter(handler *UserHandler) chi.Router {
	r := chi.NewRouter()

	// Пример маршрутов
	r.Post("/", handler.CreateUserHandler)       // Создание пользователя
	r.Get("/{id}", handler.GetUserByIDHandler)   // Получение пользователя по ID
	r.Put("/", handler.UpdateUserHandler)        // Обновление пользователя
	r.Delete("/{id}", handler.DeleteUserHandler) // Удаление пользователя

	return r
}
