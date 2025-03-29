package transaction

import (
	// "encoding/json"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func TransactionRouter(handler *TransactionHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/id/{id}", handler.GetTransactionByIDHandler)
	r.Get("/", handler.GetTransactionByUserIDHandler)
	r.Post("/", handler.CreateTransactionHandler)
	r.Put("/", handler.UpdateTransactionHandler)
	r.Delete("/id={id}", handler.DeleteTransactionHandler)

	return r
}
