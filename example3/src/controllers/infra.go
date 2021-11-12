package controllers

import (
	"api/src/banco"
	"api/src/respostas"
	"net/http"
)

func Prove(w http.ResponseWriter, r *http.Request) {
	var mongo banco.Mongo
	_, erro := mongo.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, "Sistema rodando perfeitamente")
}
