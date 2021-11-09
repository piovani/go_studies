package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizaPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
