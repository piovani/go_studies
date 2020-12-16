package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

var db *sql.DB

func pageMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem Vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	registros, err := db.Query("SELECT id, autor, titulo FROM livros")

	if err != nil {
		log.Fatal(err.Error())
	}

	var livros []Livro = make([]Livro, 0)
	for registros.Next() {
		var livro Livro
		errScan := registros.Scan(&livro.Id, &livro.Autor, &livro.Titulo)

		if errScan != nil {
			log.Fatal(errScan.Error())
			continue
		}

		livros = append(livros, livro)
	}

	registros.Close()

	encoder := json.NewEncoder(w)
	encoder.Encode(livros)
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

func configurarBancoDeDados() {
	var err error
	db, err = sql.Open(
		os.Getenv("DB_CONNECTION"),
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_DATABASE")))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configurarBancoDeDados()
	confService()
}
