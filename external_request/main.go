package main

import (
	"bytes"
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
	// post()
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

func post() string {
	body := map[string]string{
		"email": "usuario1@email.com",
		"senha": "1234",
	}

	bodyJson, _ := json.Marshal(body)

	res, _ := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(bodyJson))

	resBytes, _ := ioutil.ReadAll(res.Body)
	token := string(resBytes)

	return token
}

func update() {
	token := post()

	body := map[string]string{
		"nome":  "usuario 11a",
		"nick":  "usuario111",
		"email": "usuario11@usuario.com",
	}

	bodyJson, _ := json.Marshal(body)

	url := "http://localhost:5000/usuarios/1"

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyJson))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := client.Do(request)
	defer response.Body.Close()

	fmt.Println(response.Status)
}

func delete() {
	token := post()

	url := fmt.Sprintf("http://localhost:5000/usuarios/%d", 1)

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodDelete, url, nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := client.Do(request)
	defer response.Body.Close()

	fmt.Println(response.Status)
}
