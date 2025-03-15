package handlers

import (
	"encoding/json"
	"finance-service/cmd/transaction-service/entities"
	"finance-service/cmd/transaction-service/services"
	"fmt"
	"net/http"
)

// CreateTransaction обрабатывает запросы на создание транзакции
// @Summary Создать транзакцию
// @Description Создает новую финансовую транзакцию
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body entities.TransactionRequest true "Данные транзакции"
// @Success 201 {object} models.Transaction "Созданная транзакция"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /transactions/create [post]
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req entities.TransactionRequest

	// Декодируем тело запроса
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	// Создание транзакции через сервис
	transaction, err := services.CreateTransaction(req.UserID, req.ActionType, req.Quantity, req.TariffID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправка ответа с созданной транзакцией
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
