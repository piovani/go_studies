package router

import (
	"go_web/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	router := mux.NewRouter()
	rotas.Configurar(router)

	return router
}
