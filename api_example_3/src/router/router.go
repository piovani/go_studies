package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Vai gerar as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
