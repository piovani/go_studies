package controllers

import (
	"api/src/respostas"
	"net/http"
)

type M struct {
	By string `bson:"by"`
}

func Prove(w http.ResponseWriter, r *http.Request) {
	// var mongo banco.Mongo
	// client, erro := mongo.Conectar()
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }

	// collection := client.Database("mydb").Collection("users")
	// bson := M{By: "Teste"}

	// res, erro := collection.InserOne(context.Background(), bson)
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, "Deu ruim a inserir")
	// 	return
	// }

	// id := res.InsertedID
	// fmt.Println(id)

	respostas.JSON(w, http.StatusOK, "Sistema rodando perfeitamente")
}
