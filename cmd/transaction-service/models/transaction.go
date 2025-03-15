package models

import "time"

// Transaction представляет структуру для хранения информации о транзакции
type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ActionType  string    `json:"action_type"`
	Quantity    float64   `json:"quantity"`     // Переименовали Amount в Quantity
	TotalAmount float64   `json:"total_amount"` // Добавили поле для общей суммы
	TariffID    int       `json:"tariff_id"`
	CreatedAt   time.Time `json:"created_at"`
}
