package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/piovani/go_api/domain"
)

func Connection() *gorm.DB {

	dsn := "user:secret@tcp(127.0.0.1:33069)/database?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error during connection: %v", err)
	}

	db.AutoMigrate(domain.Livro{})

	return db
}
