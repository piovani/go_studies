package controllers

import (
	"api/src/respostas"
	"net/http"
)

func Prove(w http.ResponseWriter, r *http.Request) {
	respostas.JSON(w, http.StatusOK, "Sistema rodando perfeitamente")
}
