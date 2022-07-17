package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	//Conexão com o MySQL
	StringConexaoBanco = ""

	//Portal onde roda a API
	Porta = 5000
)

//Iniciando as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	//Estilo ParseInt
	Porta, erro := strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		//Se der um erro, trocaremos a porta para:
		Porta = 9000
		fmt.Sprintf("Porta modificada para %s", Porta)
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
		)
}