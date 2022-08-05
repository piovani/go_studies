package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasInfra = []Rota{
	{
		URI:                "/heart",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Prove,
		RequerAutenticacao: false,
	},
}
