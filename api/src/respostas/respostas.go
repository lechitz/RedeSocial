package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON : Recebe o statusCode, adiciona o status ao header, depois pega os dados genéricos e transforma pra JSON. É respondido um JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

//Erro: retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		//O método Error dentro do tipo erro, é um método que retorna a mensagem de erro
		Erro: erro.Error(),
	})
}