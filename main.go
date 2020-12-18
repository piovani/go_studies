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

	registros := db.QueryRow("SELECT id, autor, titulo FROM livros WHERE id = ?", id)
	var livro Livro

	err := registros.Scan(&livro.Id, &livro.Autor, &livro.Titulo)

	if err != nil {
		log.Println("Livro não encontrado, erro: " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(livro)
	w.WriteHeader(http.StatusNotFound)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	body, _ := ioutil.ReadAll(r.Body)

	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)

	if len(novoLivro.Autor) == 0 || len(novoLivro.Titulo) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, _ := db.Exec("INSERT INTO livros (autor, titulo) VALUES (?, ?)", novoLivro.Autor, novoLivro.Titulo)

	novoId, _ := result.LastInsertId()
	novoLivro.Id = int(novoId)

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

	registro := db.QueryRow("SELECT id, autor, titulo FROM livros WHERE id = ?", id)
	var livro Livro

	err4 := registro.Scan(&livro.Id, &livro.Autor, &livro.Titulo)

	if err4 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	LivroModificado.Id = id

	_, err5 := db.Exec("UPDATE livros SET autor = ?, titulo = ? WHERE id = ?",
		LivroModificado.Autor,
		LivroModificado.Titulo,
		id)

	if err5 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(LivroModificado)
}

func excluirLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	registro := db.QueryRow("SELECT id FROM livros WHERE id = ?", id)
	var idLivro int
	registro.Scan(idLivro)

	db.Exec("DELETE FROM livros WHERE id = ?", id)

	w.WriteHeader(http.StatusNoContent)
	return
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
