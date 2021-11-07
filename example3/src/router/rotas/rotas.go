package rotas

import (
	"api/middlewares"
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
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}

func getRotas() []Rota {
	var rotas []Rota

	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotasInfra...)
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	return rotas
}
