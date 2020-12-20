package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/piovani/go_api/domain"
)

func Connection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error during Load dotenv: %v", err)
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE")))

	if err != nil {
		log.Fatalf("Error during connection: %v", err)
	}

	// defer db.Close()

	db.AutoMigrate(domain.Livro{})

	return db
}
