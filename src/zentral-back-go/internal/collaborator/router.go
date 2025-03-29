package collaborator

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func CollaboratorRouter(handler *CollaboratorHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateCollaboratorHandler)
	r.Get("/{id}", handler.GetCollaboratorByIDHandler)
	r.Put("/", handler.UpdateCollaboratorHandler)
	r.Delete("/{id}", handler.DeleteCollaboratorHandler)

	return r
}
