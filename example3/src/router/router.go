package router

import "github.com/gorilla/mux"

// Vai gerar as rotas configuradas
func Generate() *mux.Router {
	return mux.NewRouter()
}
