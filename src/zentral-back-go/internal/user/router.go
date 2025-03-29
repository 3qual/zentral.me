package user

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(handler *UserHandler) chi.Router {
	r := chi.NewRouter()

	// Пример маршрутов
	r.Post("/", handler.CreateUserHandler)       // Создание пользователя
	r.Get("/{id}", handler.GetUserByIDHandler)   // Получение пользователя по ID
	r.Put("/", handler.UpdateUserHandler)        // Обновление пользователя
	r.Delete("/{id}", handler.DeleteUserHandler) // Удаление пользователя

	return r
}
