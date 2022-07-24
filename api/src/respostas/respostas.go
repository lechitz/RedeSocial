package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON : Recebe o statusCode, adiciona o status ao header, depois pega os dados genéricos e transforma pra JSON. É respondido um JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	//Tipo do conteúdo da resposta vai ser um JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

//Erro : retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		//O método Error dentro do tipo erro, é um método que retorna a mensagem de erro
		Erro: erro.Error(),
	})
}