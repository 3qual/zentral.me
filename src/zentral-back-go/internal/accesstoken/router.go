package accesstoken

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func AccessTokenRouter(handler *AccessTokenHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateAccessTokenHandler)
	r.Get("/{id}", handler.GetAccessTokenByIDHandler)
	r.Put("/", handler.UpdateAccessTokenHandler)
	r.Delete("/{id}", handler.DeleteAccessTokenHandler)

	return r
}
