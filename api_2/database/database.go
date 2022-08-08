package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (database *Database, err error) {
	client, err := getConnection()
	if err != nil {
		return database, err
	}

	return &Database{
		Client: client,
	}, nil
}

func getConnection() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open("mysql", dns)
	if err != nil {
		return db, err
	}

	return db, err
}
