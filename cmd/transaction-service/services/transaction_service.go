package services

import (
	"encoding/json"
	"finance-service/cmd/transaction-service/config"
	"finance-service/cmd/transaction-service/db"
	"finance-service/cmd/transaction-service/entities"
	"finance-service/cmd/transaction-service/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetTariffPrice - получает цену тарифа по его ID
func GetTariffPrice(tariffID int) (float64, error) {
	url := fmt.Sprintf("%s/tariffs/get?id=%d", config.AppConfig.TariffServiceURL, tariffID)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error requesting tariff price: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("tariff service returned status: %d", resp.StatusCode)
	}

	var tariff entities.TariffResponse
	if err := json.NewDecoder(resp.Body).Decode(&tariff); err != nil {
		return 0, fmt.Errorf("error decoding tariff response: %v", err)
	}

	return tariff.Price, nil
}

// CreateTransaction создает новую транзакцию в базе данных
func CreateTransaction(userID int, actionType string, quantity float64, tariffID int) (models.Transaction, error) {
	// Получаем цену тарифа из другого сервиса
	price, err := GetTariffPrice(tariffID)
	if err != nil {
		log.Printf("Error fetching tariff price: %v", err)
		return models.Transaction{}, err
	}

	// Рассчитываем общую сумму
	totalAmount := price * quantity

	// Формируем запрос для вставки новой транзакции
	query := `INSERT INTO transactions (user_id, action_type, quantity, total_amount, tariff_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, action_type, quantity, total_amount, tariff_id, created_at`

	var transaction models.Transaction
	err = db.TransactionDB.QueryRow(query, userID, actionType, quantity, totalAmount, tariffID, time.Now()).Scan(&transaction.ID, &transaction.UserID, &transaction.ActionType, &transaction.Quantity, &transaction.TotalAmount, &transaction.TariffID, &transaction.CreatedAt)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return models.Transaction{}, err
	}

	// Возвращаем созданную транзакцию
	return transaction, nil
}
