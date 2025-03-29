package folder

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FolderRouter(handler *FolderHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateFolderHandler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		folders, err := handler.service.GetAllFolders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(folders)
	})
	r.Get("/{id}", handler.GetFolderByIDHandler)
	r.Put("/", handler.UpdateFolderHandler)
	r.Delete("/{id}", handler.DeleteFolderHandler)

	return r
}
