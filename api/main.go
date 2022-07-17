package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API")
	r := router.Gerar()

	//Subindo o servidor em uma porta
	log.Fatal(http.ListenAndServe(":5000", r))
}
