package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Rota representa todas as rotas da API
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request) //Todas as funções dentro da rota terão esse formato
	RequerAutenticacao bool //Ver se a rota precisa que o usuario esteja logado ou não
}

//Retorna o Router com todas as rotas já configuradas
func Configurar(r *mux.Router) *mux.Router {
	//Chamo os routers vazios
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	//Crio o range para setar todos os routers (URI, função e método)
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}