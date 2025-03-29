package transaction

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func TransactionRouter(handler *TransactionHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateTransactionHandler)
	r.Get("/{id}", handler.GetTransactionByIDHandler)
	r.Put("/", handler.UpdateTransactionHandler)
	r.Delete("/{id}", handler.DeleteTransactionHandler)

	return r
}
