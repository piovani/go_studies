package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/piovani/go_api/database"
)

func CheckFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func StartConfig() error {
	return godotenv.Load()
}

func NewDatabase() (*database.Database, error) {
	return database.NewDatabase()
}
