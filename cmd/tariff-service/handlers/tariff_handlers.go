// internal/handlers/tariff.go

package handlers

import (
	"encoding/json"
	"finance-service/cmd/tariff-service/models"
	"finance-service/cmd/tariff-service/services"
	"fmt"
	"net/http"
	"strconv"
)

// GetAllTariffs возвращает список всех тарифов
// @Summary Получить все тарифы
// @Description Возвращает список всех тарифов в системе
// @Tags tariffs
// @Produce json
// @Success 200 {array} models.Tariff
// @Router /tariffs [get]
func GetAllTariffs(w http.ResponseWriter, r *http.Request) {
	// Получаем все тарифы
	tariffs, err := services.GetAllTariffs()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching tariffs: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем список тарифов в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tariffs)
}

// GetTariffByID получает тариф по ID
// @Summary Получить тариф по ID
// @Description Возвращает данные тарифа по его идентификатору
// @Tags tariffs
// @Produce json
// @Param id query int true "ID тарифа"
// @Success 200 {object} models.Tariff
// @Router /tariffs/get [get]
func GetTariffByID(w http.ResponseWriter, r *http.Request) {
	// Получаем ID тарифа из URL
	// Например, URL: /tariffs?id=1
	idStr := r.URL.Query().Get("id") // id приходит как строка

	// Преобразуем строку в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	// Получаем тариф по ID
	tariff, err := services.GetTariffByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching tariff: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем тариф в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tariff)
}

// CreateTariff создаёт новый тариф
// @Summary Создать тариф
// @Description Добавляет новый тариф в систему
// @Tags tariffs
// @Accept json
// @Produce json
// @Param tariff body models.TariffRequest true "Данные нового тарифа"
// @Success 201 {object} models.TariffRequest
// @Router /tariffs/create [post]
func CreateTariff(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем тело запроса в структуру Tariff
	var tariff models.TariffRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tariff)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Создаем тариф
	err = services.CreateTariff(tariff)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating tariff: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tariff)
}
