package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Domain struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdateAt  time.Time `json:"updated_at" gorm:"type:datetime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:datetime"`
}

func (domain *Domain) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during obs creating: %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during obs ID create: %v", err)
	}

	return nil
}
