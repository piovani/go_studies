package main

import (
	"log"

	"github.com/piovani/go_api/database"
	"github.com/piovani/go_api/domain"
)

func main() {
	db := database.Connection()

	livro := domain.Livro{
		Titulo: "teste titulo",
		Autor:  "teste autor",
	}

	err := db.Create(livro).Error

	if err != nil {
		log.Fatalf("Error during create livro: %v", err)
	}

}
