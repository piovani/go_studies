package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Livro struct {
	Id     int
	Titulo string
	Autor  string
}

var Livros []Livro = []Livro{
	Livro{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "José de Alencar",
	},
	Livro{
		Id:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	},
	Livro{
		Id:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func pageMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem Vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func confRoutes() {
	http.HandleFunc("/", pageMain)
	http.HandleFunc("/livros", listarLivros)
}

func confService() {
	port := "1337"

	confRoutes()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	confService()
	fmt.Println("Servidor esta rodando na porta")
}
