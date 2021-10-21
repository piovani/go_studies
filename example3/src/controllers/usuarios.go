package controllers

import (
	"api/modelos"
	"api/modelos"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()




	w.Write([]byte("Criando Usuário!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
