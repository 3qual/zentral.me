package foldertransaction

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// FolderTransactionHandler структура для обработчика связи папки и транзакции
type FolderTransactionHandler struct {
	service FolderTransactionService
}

// NewFolderTransactionHandler конструктор для создания нового обработчика
func NewFolderTransactionHandler(service FolderTransactionService) *FolderTransactionHandler {
	return &FolderTransactionHandler{service: service}
}

// CreateFolderTransactionHandler создаёт новую связь
func (h *FolderTransactionHandler) CreateFolderTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var folderTransaction FolderTransaction
	if err := json.NewDecoder(r.Body).Decode(&folderTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateFolderTransaction(&folderTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(folderTransaction)
}

// GetFolderTransactionByIDHandler получает связь по ID
func (h *FolderTransactionHandler) GetFolderTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	folderTransaction, err := h.service.GetFolderTransactionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folderTransaction)
}

// GetFolderTransactionsByFolderIDHandler получает все связи по ID папки
func (h *FolderTransactionHandler) GetFolderTransactionsByFolderIDHandler(w http.ResponseWriter, r *http.Request) {
	folderID := chi.URLParam(r, "folder_id")
	folderTransactions, err := h.service.GetFolderTransactionsByFolderID(uuid.Must(uuid.Parse(folderID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folderTransactions)
}

// GetFolderTransactionsByTransactionIDHandler получает все связи по ID транзакции
func (h *FolderTransactionHandler) GetFolderTransactionsByTransactionIDHandler(w http.ResponseWriter, r *http.Request) {
	transactionID := chi.URLParam(r, "transaction_id")
	folderTransactions, err := h.service.GetFolderTransactionsByTransactionID(uuid.Must(uuid.Parse(transactionID)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folderTransactions)
}

// DeleteFolderTransactionHandler удаляет связь
func (h *FolderTransactionHandler) DeleteFolderTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	folderTransaction, err := h.service.GetFolderTransactionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteFolderTransaction(folderTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
