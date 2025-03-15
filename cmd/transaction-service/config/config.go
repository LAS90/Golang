package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TransactionDBDSN string
	TariffServiceURL string
}

var AppConfig Config

func LoadConfig() {
	// Загрузка файла .env
	err := godotenv.Load("cmd/transaction-service/config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Инициализация структуры конфигурации
	AppConfig = Config{
		TransactionDBDSN: os.Getenv("TRANSACTION_DB_DSN"),
		TariffServiceURL: os.Getenv("TARIFF_SERVICE_URL"),
	}

	// Проверка на обязательные переменные
	if AppConfig.TransactionDBDSN == "" {
		log.Fatal("TRANSACTION_DB_DSN is not set")
	}
	if AppConfig.TariffServiceURL == "" {
		log.Fatal("TARIFF_SERVICE_URL is not set")
	}
}
