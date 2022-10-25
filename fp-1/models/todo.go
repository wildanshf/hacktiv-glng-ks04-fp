package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Description string `json:"description"`
}

type TodoPayload struct {
	Description string `json:"description"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
