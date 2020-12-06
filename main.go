package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func listarLivros(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func buscarLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, livro := range Livros {
		if livro.Id == id {
			json.NewEncoder(w).Encode(livro)
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
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

func modificarLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	corpo, err2 := ioutil.ReadAll(r.Body)

	var LivroModificado Livro
	err3 := json.Unmarshal(corpo, &LivroModificado)

	if err != nil || err2 != nil || err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indiceDoLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceDoLivro = indice
			break
		}
	}

	if indiceDoLivro < 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	LivroModificado.Id = id
	Livros[indiceDoLivro] = LivroModificado

	json.NewEncoder(w).Encode(LivroModificado)
}

func excluirLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	indeceLivro := -1
	for indece, livro := range Livros {
		if livro.Id == id {
			indeceLivro = indece
			break
		}
	}

	if indeceLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	ladoEsquerdo := Livros[0:indeceLivro]
	ladoDireito := Livros[indeceLivro+1 : len(Livros)]

	Livros = append(ladoEsquerdo, ladoDireito...)

	w.WriteHeader(http.StatusNoContent)
}

func confRoutes(router *mux.Router) {
	router.HandleFunc("/", pageMain).Methods("GET")
	router.HandleFunc("/livros", listarLivros).Methods("GET")
	router.HandleFunc("/livros/{id}", buscarLivro).Methods("GET")
	router.HandleFunc("/livros", cadastrarLivro).Methods("POST")
	router.HandleFunc("/livros/{id}", modificarLivro).Methods("PUT")
	router.HandleFunc("/livros/{id}", excluirLivro).Methods("DELETE")
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func confService() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	port := "1337"

	confRoutes(router)
	fmt.Println("Servidor esta rodando")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	confService()
}
