package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// TariffDB - глобальная переменная для хранения подключения к БД
var TariffDB *sql.DB

// InitTariffDB - инициализация подключения к БД тарифов
func InitTariffDB() (*sql.DB, error) {
	// Читаем строку подключения из переменных окружения
	connStr := os.Getenv("TARIFF_DB_DSN")
	if connStr == "" {
		log.Fatal("TARIFF_DB_DSN is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	TariffDB = db
	fmt.Println("Successfully connected to the tariff database.")
	return db, nil
}

// CloseTariffDB - закрывает соединение с БД
func CloseTariffDB() {
	if TariffDB != nil {
		if err := TariffDB.Close(); err != nil {
			log.Printf("Error closing tariff database connection: %v", err)
		}
	}
}
