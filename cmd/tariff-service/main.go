package main

import (
	"log"
	"net/http"
	"os"

	"finance-service/cmd/tariff-service/db"
	_ "finance-service/cmd/tariff-service/docs"
	"finance-service/cmd/tariff-service/handlers"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Tariff Service API
// @version 1.0
// @description API для управления тарифами
// @host localhost:8081
// @BasePath /
func main() {
	// Загружаем .env (если файл есть)
	err1 := godotenv.Load("cmd/tariff-service/config.env")
	if err1 != nil {
		log.Fatalf("Error loading .env file: %v", err1)
	}

	// Проверяем, загружены ли переменные окружения
	if os.Getenv("TARIFF_DB_DSN") == "" {
		log.Fatal("TARIFF_DB_DSN is not set")
	}

	// Инициализация подключения к базе данных тарифов
	_, err := db.InitTariffDB()
	if err != nil {
		log.Fatalf("Error initializing tariff database: %v", err)
	}
	defer db.CloseTariffDB() // Закрытие соединения при завершении работы

	// Роутинг
	http.HandleFunc("/tariffs", handlers.GetAllTariffs)
	http.HandleFunc("/tariffs/create", handlers.CreateTariff)
	http.HandleFunc("/tariffs/", handlers.GetTariffByID)

	// Подключение Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Запуск HTTP-сервера
	log.Println("Tariff service running on port 8081...")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
