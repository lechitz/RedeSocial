package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

//Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	//r é do tipo router, então pode ser passado para a função configurar
	r :=  mux.NewRouter()
	//a função Configurar devolve o mux.Router
	return rotas.Configurar(r)
}