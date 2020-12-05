package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
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

func rotearLivros(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarLivros(w, r)
	case "POST":
		cadastrarLivro(w, r)
	}
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(encoder.Encode("DEU RUIM IRMAO"))
	}

	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1
	Livros = append(Livros, novoLivro)

	encoder.Encode(novoLivro)
}

func buscarLivro(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")

	if len(partes) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(partes[2])

	for _, livro := range Livros {
		if livro.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(livro)
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func confRoutes() {
	http.HandleFunc("/", pageMain)
	http.HandleFunc("/livros", rotearLivros)
	http.HandleFunc("/livros/", buscarLivro)
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
