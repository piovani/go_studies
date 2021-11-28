package models

import (
	"time"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description"`
	MediumPrice string    `json:"medium_price"`
	Author      string    `json:"author"`
	ImageURL    string    `json:"img_url"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   time.Time `json:"deleted"`
}
