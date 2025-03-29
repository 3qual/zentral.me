package foldertransaction

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func FolderTransactionRouter(handler *FolderTransactionHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateFolderTransactionHandler)
	r.Get("/{id}", handler.GetFolderTransactionByIDHandler)
	r.Delete("/{id}", handler.DeleteFolderTransactionHandler)

	return r
}
