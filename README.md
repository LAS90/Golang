# Finance Service

Это приложение состоит из двух микросервисов: **Tariff Service** и **Transaction Service**. Эти сервисы работают вместе для управления тарифами и транзакциями.

## Содержание

- [Обзор](#Обзор)
- [Установка](#Установка)
- [Запуск](#Запуск)
- [API](#API)
- [Структура базы данных](#Структура-базы-данных)

## Обзор

### Tariff Service
Сервис для управления тарифами. Он позволяет добавлять, изменять и получать тарифы. Сервис предоставляет API для получения всех тарифов и информации по тарифу по его ID.

### Transaction Service
Сервис для управления транзакциями пользователей. Этот сервис создает транзакции, рассчитывает общую сумму на основе тарифов и сохраняет данные о транзакциях в базе данных. Он также взаимодействует с сервисом тарифов для получения цены по ID тарифа.

## Установка

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/LAS90/Golang.git
    cd Golang
    ```

2. Установите зависимости:

    ```bash
    go mod tidy
    ```

3. Создайте файлы конфигурации для каждого сервиса:
    - **config.env** — должен содержать строки подключения к базе данных для каждого сервиса и другие конфигурационные параметры.

    Пример для **Transaction Service**:
    ```ini
    TRANSACTION_DB_DSN=your_db_connection_string
    ```

    Пример для **Tariff Service**:
    ```ini
    TARIFF_DB_DSN=your_db_connection_string
    ```

## Запуск

### Запуск Tariff Service

1. Запустите сервис:

    ```bash
    go run cmd/tariff-service/main.go
    ```

2. Сервис будет доступен по адресу `http://localhost:8081`.

### Запуск Transaction Service

1. Запустите сервис:

    ```bash
    go run cmd/transaction-service/main.go
    ```

2. Сервис будет доступен по адресу `http://localhost:8082`.

## API

### Tariff Service

- **GET /tariffs** — Получить все тарифы.
- **GET /tariffs/get?id={id}** — Получить тариф по ID.
- **POST /tariffs/create** — Создать новый тариф.

### Transaction Service

- **POST /transactions/create** — Создать новую транзакцию.

## Структура базы данных

### Tariff Service

База данных для Tariff Service хранит тарифы. Структура таблицы:

```sql
CREATE TABLE tariffs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
### Transaction Service

База данных для Transaction Service хранит транзакции. Структура таблицы:

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    action_type VARCHAR(255) NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    tariff_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tariff_id) REFERENCES tariffs(id)
);
```
