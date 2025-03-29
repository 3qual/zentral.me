package image

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func ImageRouter(handler *ImageHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateImageHandler)
	r.Get("/{id}", handler.GetImageByIDHandler)
	r.Delete("/{id}", handler.DeleteImageHandler)

	return r
}
