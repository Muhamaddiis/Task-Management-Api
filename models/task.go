package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"unique;not null"`
	Description string         `json:"description"`
	Status      string         `json:"status" gorm:"default:pending"`
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TaskRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"oneof=pending in_progress completed"`
	DueDate     time.Time `json:"due_date" validate:"required"`
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskQueryParams struct {
	Status  string `json:"status"`
	DueDate string `json:"due_date"`
	Search  string `json:"search"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}