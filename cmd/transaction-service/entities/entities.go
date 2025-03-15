// entities/tariff.go
package entities

// TariffResponse представляет структуру ответа с тарифом
type TariffResponse struct {
	Price float64 `json:"price"`
}

// TransactionRequest используется для запроса на создание транзакции
type TransactionRequest struct {
	UserID     int     `json:"user_id" example:"1"`
	ActionType string  `json:"action_type" example:"purchase"`
	Quantity   float64 `json:"quantity" example:"99.99"` // Переименовали Amount в Quantity
	TariffID   int     `json:"tariff_id" example:"2"`
}
