package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:secret@tcp(127.0.0.1:33069)/database")

	if err != nil {
		panic(err.Error())
	}

	// _, errCreate := db.Exec(
	// 	"CREATE TABLE livros (" +
	// 		"id INT NOT NULL AUTO_INCREMENT," +
	// 		"autor VARCHAR(255) NOT NULL," +
	// 		"titulo VARCHAR(255) NOT NULL," +
	// 		"PRIMARY KEY (id)" +
	// 		")")

	// if errCreate != nil {
	// 	log.Fatal(errCreate.Error())
	// }

	// _, errInsert := db.Exec("INSERT INTO livros (autor, titulo) VALUES ('José de Alencar', 'O Guarani')")
	_, errInsert := db.Exec("INSERT INTO livros (autor, titulo) VALUES ('Cazuza', 'Viriato Correia')")
	_, errInsert2 := db.Exec("INSERT INTO livros (autor, titulo) VALUES ('Dom Casmurro', 'Machado de Assis')")

	if (errInsert != nil) || (errInsert2 != nil) {
		log.Fatal(errInsert.Error())
	}
}
