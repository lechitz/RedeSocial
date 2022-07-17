package controllers

import "net/http"

//Contém as funções de usuários

//response deveria ser request

//Criando um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando usuário."))
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