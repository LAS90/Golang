package services

import (
	"finance-service/cmd/tariff-service/db"
	"finance-service/cmd/tariff-service/models"
	"fmt"
	"time"
)

// GetAllTariffs - функция для получения всех тарифов
func GetAllTariffs() ([]models.Tariff, error) {
	rows, err := db.TariffDB.Query("SELECT id, action_type, price, valid_from, valid_until FROM tariffs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tariffs []models.Tariff
	for rows.Next() {
		var tariff models.Tariff
		if err := rows.Scan(&tariff.ID, &tariff.ActionType, &tariff.Price, &tariff.ValidFrom, &tariff.ValidUntil); err != nil {
			return nil, err
		}
		tariffs = append(tariffs, tariff)
	}
	return tariffs, nil
}

// CreateTariff создает новый тариф в базе данных, принимая модель запроса
func CreateTariff(tariffRequest models.TariffRequest) error {
	// Формируем даты valid_from и valid_until
	validFrom := time.Now().Format("2006-01-02")                   // Текущая дата
	validUntil := time.Now().AddDate(0, 1, 0).Format("2006-01-02") // +1 месяц

	// Вставляем новый тариф в базу данных
	_, err := db.TariffDB.Exec("INSERT INTO tariffs (action_type, price, valid_from, valid_until) VALUES ($1, $2, $3, $4)",
		tariffRequest.ActionType, tariffRequest.Price, validFrom, validUntil)
	if err != nil {
		return err
	}
	fmt.Printf("New tariff created with ActionType: %s, Price: %f, ValidFrom: %s, ValidUntil: %s\n",
		tariffRequest.ActionType, tariffRequest.Price, validFrom, validUntil)
	return nil
}

// GetTariffByID - функция для получения тарифа по ID
func GetTariffByID(id int) (models.Tariff, error) {
	var tariff models.Tariff
	err := db.TariffDB.QueryRow("SELECT id, action_type, price, valid_from, valid_until FROM tariffs WHERE id = $1", id).
		Scan(&tariff.ID, &tariff.ActionType, &tariff.Price, &tariff.ValidFrom, &tariff.ValidUntil)
	if err != nil {
		return tariff, err
	}
	return tariff, nil
}
