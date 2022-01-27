package main

import (
	"fmt"
	"go_web/src/router"
	"log"
	"net/http"

	"go_web/src/utils"
)

func main() {
	fmt.Println("Rodando")

	utils.CarregarTemplates()

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", r))
}
