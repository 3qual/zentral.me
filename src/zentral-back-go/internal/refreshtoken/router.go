package refreshtoken

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func RefreshTokenRouter(handler *RefreshTokenHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateRefreshTokenHandler)
	r.Get("/{id}", handler.GetRefreshTokenByIDHandler)
	r.Put("/", handler.UpdateRefreshTokenHandler)
	r.Delete("/{id}", handler.DeleteRefreshTokenHandler)

	return r
}
