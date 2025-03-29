package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// TransactionHandler структура для обработчика транзакции
type TransactionHandler struct {
	service TransactionService
}

// NewTransactionHandler конструктор для создания нового обработчика
func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// CreateTransactionHandler создаёт новую транзакцию
func (h *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTransaction(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// GetTransactionByIDHandler получает транзакцию по ID
func (h *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	transaction, err := h.service.GetTransactionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// GetTransactionByUserIDHandler получает все транзакции по ID пользователя с пагинацией
func (h *TransactionHandler) GetTransactionByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	// page := r.URL.Query().Get("page")   // Параметр страницы
	// limit := r.URL.Query().Get("limit") // Параметр лимита

	// Если параметры пустые, присваиваем им значения по умолчанию
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	// Преобразуем параметры в числа
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	// Получаем транзакции с пагинацией
	transactions, err := h.service.GetTransactionByUserID(uuid.Must(uuid.Parse(userID)), pageInt, limitInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// UpdateTransactionHandler обновляет данные транзакции
func (h *TransactionHandler) UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateTransaction(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// DeleteTransactionHandler удаляет транзакцию
func (h *TransactionHandler) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	transaction, err := h.service.GetTransactionByID(uuid.Must(uuid.Parse(id)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.DeleteTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
