package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := "user"
	password := "secret"
	database := "database"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:33069)/%s", user, password, database))

	if err != nil {
		panic(err.Error())
	}

	_, errDrop := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s.livros", database))

	if errDrop != nil {
		panic(errDrop.Error())
	}

	_, errCreate := db.Exec(
		"CREATE TABLE livros (" +
			"id INT NOT NULL AUTO_INCREMENT," +
			"autor VARCHAR(255) NOT NULL," +
			"titulo VARCHAR(255) NOT NULL," +
			"PRIMARY KEY (id)" +
			")")

	if errCreate != nil {
		log.Fatal(errCreate.Error())
	}

	_, errInsert := db.Exec("INSERT INTO livros (autor, titulo) VALUES " +
		"('José de Alencar', 'O Guarani')," +
		"('Cazuza', 'Viriato Correia')," +
		"('Dom Casmurro', 'Machado de Assis')")

	if errInsert != nil {
		log.Fatal(errInsert.Error())
	}
}
