package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Body struct {
	Navegacao struct {
		ID                string `json:"ID"`
		Titulo            string `json:"Titulo"`
		Caminho           string `json:"Caminho"`
		Servico           int64  `json:"Servico"`
		ParametrosServico string `json:"ParametrosServico"`
		Link              string `json:"Link"`
		Descricao         string `json:"Descricao"`
		Raiz              string `json:"Raiz"`
		OcultarDoMenu     bool   `json:"OcultarDoMenu"`
		Descendentes      []struct {
			ID                string `json:"ID"`
			Titulo            string `json:"Titulo"`
			Caminho           string `json:"Caminho"`
			Servico           int64  `json:"Servico"`
			ParametrosServico string `json:"ParametrosServico"`
			Link              string `json:"Link"`
			Descricao         string `json:"Descricao"`
			Raiz              bool   `json:"Raiz"`
			OcultarDoMenu     bool   `json:"OcultarDoMenu"`
			Descendentes      string `json:"Descendentes"`
		} `json:"Descendentes"`
	} `json:"navegacao"`
	View     string `json:"view"`
	Conteudo []struct {
		Data      time.Time `json:"DataReuniaoCopom"`
		Vues      string    `json:"Vies"`
		MetaSelic float64   `json:"MetaSelic"`
	} `json:"conteudo"`
}

func main() {
	// get()
	post()
	// update()
	// delete()
}

func get() {
	res, _ := http.Get("https://www.bcb.gov.br/api/conteudo/sitebcb/indicadortaxaselic/ultima")
	defer res.Body.Close()

	var body Body
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &body)

	fmt.Println(body.Navegacao.Descendentes[0].ID)
}

func post() {
	res, _ := http.Get("https://www.bcb.gov.br/api/conteudo/sitebcb/indicadortaxaselic/ultima")
	defer res.Body.Close()

	var body Body
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &body)

	fmt.Println(body.Navegacao.Descendentes[0].ID)
}
