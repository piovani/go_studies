package database

import (
	"database/sql"
	"fmt"
	"log"
)

func SaveMessage(message string) {
	db := getConnectionDB()
	statement, err := db.Prepare(
		"INSERT INTO messages (message) VALUES (?)",
	)

	if err != nil {
		log.Fatal("Erro: ", err)
	}
	defer statement.Close()

	res, err := statement.Exec(message)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	fmt.Println(lastID)
}

func getConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "localhost:3306")
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal("Erro: ", err)
	}

	return db
}
