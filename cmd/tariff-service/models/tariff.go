package models

// Tariff - модель для хранения информации о тарифах
type Tariff struct {
	ID         int     `json:"id"`
	ActionType string  `json:"action_type" example:"Обработка документа"`  // Тип действия (например, "Обработка документа")
	Price      float64 `json:"price" example:"100.50"`                     // Цена тарифа
	ValidFrom  string  `json:"valid_from,omitempty" example:"2025-03-14"`  // Начало действия тарифа
	ValidUntil string  `json:"valid_until,omitempty" example:"2025-12-31"` // Окончание действия тарифа
}

// TariffRequest - модель для создания или обновления тарифа, без ID и дат
type TariffRequest struct {
	ActionType string  `json:"action_type" example:"Обработка документа"` // Тип действия (например, "Обработка документа")
	Price      float64 `json:"price" example:"10.50"`                     // Цена тарифа
	ValidFrom  string  `json:"valid_from" example:"2025-03-14"`           // Начало действия тарифа
	ValidUntil string  `json:"valid_until" example:"2025-12-31"`          // Окончание действия тарифа
}
