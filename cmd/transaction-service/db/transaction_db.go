package db

import (
	"database/sql"
	"finance-service/cmd/transaction-service/config"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// TransactionDB - глобальная переменная для хранения подключения к базе данных транзакций
var TransactionDB *sql.DB

// InitTransactionDB - инициализация подключения к базе данных транзакций
func InitTransactionDB() (*sql.DB, error) {
	// Используем строку подключения из конфигурации
	connStr := config.AppConfig.TransactionDBDSN

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

	TransactionDB = db
	fmt.Println("Successfully connected to the transaction database.")
	return db, nil
}

// CloseTransactionDB - закрывает соединение с базой данных транзакций
func CloseTransactionDB() {
	if err := TransactionDB.Close(); err != nil {
		log.Printf("Error closing transaction database connection: %v", err)
	}
}
