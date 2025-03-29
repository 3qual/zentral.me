package session

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func SessionRouter(handler *SessionHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateSessionHandler)
	r.Get("/{id}", handler.GetSessionByIDHandler)
	r.Put("/", handler.UpdateSessionHandler)
	r.Delete("/{id}", handler.DeleteSessionHandler)

	return r
}
