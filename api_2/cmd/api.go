package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/piovani/go_api/database"
)

type Api struct {
	DB *database.Database
}

func NewApi() *Api {
	err := StartConfig()
	CheckFatalError(err)

	db, err := NewDatabase()
	CheckFatalError(err)

	return &Api{
		DB: db,
	}
}

func (a *Api) Start() {
	a.getRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), nil)
	CheckFatalError(err)
}

func (a *Api) getRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/create", createLivro)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createLivro(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	for key, _ := range r.Form {
		fmt.Println(key)
		//LOG: {"test": "that"}
		//     err := json.Unmarshal([]byte(key), &t)
		//     if err != nil {
		//         log.Println(err.Error())
		// //     }
	}
	// log.Println(t.Test)
	//LOG: that

	// titulo := r.Body.Read("titulo")

	// fmt.Println(titulo)

	// fmt.Fprintf(w, "Welcome to the HomePage!")
	// fmt.Println("Endpoint Hit: homePage")
}
