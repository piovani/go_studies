package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasInfra = []Rota{
	{
		URI:                "/prove",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Prove,
		RequerAutenticacao: false,
	},
}
