package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura base das rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {

	rotas := getRotas()

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

func getRotas() []Rota {
	var rotas []Rota

	for _, rota := range rotasUsuarios {
		rotas = append(rotas, rota)
	}

	for _, rota := range rotasInfra {
		rotas = append(rotas, rota)
	}

	rotas = append(rotas, rotaLogin)

	return rotas
}
