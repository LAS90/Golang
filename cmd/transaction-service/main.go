package main

import (
	"log"
	"net/http"

	"finance-service/cmd/transaction-service/config"
	"finance-service/cmd/transaction-service/db"
	_ "finance-service/cmd/transaction-service/docs"
	"finance-service/cmd/transaction-service/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Transaction Service API
// @version 1.0
// @description API для управления транзакциями
// @host localhost:8082
// @BasePath /
func main() {
	// Загружаем конфигурацию
	config.LoadConfig()

	// Инициализация подключения к базе данных транзакций
	_, err := db.InitTransactionDB()
	if err != nil {
		log.Fatalf("Error initializing transaction database: %v", err)
	}
	defer db.CloseTransactionDB() // Закрытие соединения при завершении работы

	// Роутинг
	http.HandleFunc("/transactions/create", handlers.CreateTransaction)

	// Подключение Swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Запуск HTTP-сервера
	log.Println("Transaction service running on port 8082...")
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
