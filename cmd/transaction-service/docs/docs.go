// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions/create": {
            "post": {
                "description": "Создает новую финансовую транзакцию",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Создать транзакцию",
                "parameters": [
                    {
                        "description": "Данные транзакции",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Созданная транзакция",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.TransactionRequest": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "string",
                    "example": "purchase"
                },
                "quantity": {
                    "description": "Переименовали Amount в Quantity",
                    "type": "number",
                    "example": 99.99
                },
                "tariff_id": {
                    "type": "integer",
                    "example": 2
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "description": "Переименовали Amount в Quantity",
                    "type": "number"
                },
                "tariff_id": {
                    "type": "integer"
                },
                "total_amount": {
                    "description": "Добавили поле для общей суммы",
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Transaction Service API",
	Description:      "API для управления транзакциями",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
