package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Contém as funções de usuários

//response deveria ser request

//Criando um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//Criando o repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	//Chamando o usuário que tá vindo na requisição
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

//Buscando por todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os usuários."))
}

//Buscando usuário específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuário."))
}

//Atualizando usuário específico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando usuário."))
}

//Deletando usuário específico
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando usuário."))
}