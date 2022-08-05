package models

import (
	"time"
)

type Book struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	MediumPrice float64   `json:"medium_price" gorm:"medium_price"`
	Author      string    `json:"author" gorm:"author"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"deleted_at"`
}
