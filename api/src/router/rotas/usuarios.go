package rotas

import (
	"api/src/controllers"
	"net/http"
)

//Todas as rotas de usuários dentro da API
 var rotasUsuarios = []Rota {
	 {
		 //Método POST
		 URI: "/usuarios",
		 Metodo: http.MethodPost,
		 Funcao: controllers.CriarUsuario,
		 RequerAutenticacao: false,
	 },
	 {
		 //Método GET para todos os usuários: usuariosId
		 URI: "/usuarios",
		 Metodo: http.MethodGet,
		 Funcao: controllers.BuscarUsuarios,
		 RequerAutenticacao: false,
	 },
	 {
		 //Método GET pra um usuário específico: usuarioId
		 URI: "/usuarios/{usuarioId}",
		 Metodo: http.MethodGet,
		 Funcao: controllers.BuscarUsuario,
		 RequerAutenticacao: false,
	 },
	 {
		 //Método PUT para um usuário específico: usuarioId
		 URI: "/usuarios/{usuarioId}",
		 Metodo: http.MethodPut,
		 Funcao: controllers.AtualizarUsuario,
		 RequerAutenticacao: false,
	 },
	 {
		 //Método DELETE para um usuário específico: usuarioId
		 URI: "/usuarios/{usuarioId}",
		 Metodo: http.MethodDelete,
		 Funcao: controllers.DeletarUsuario,
		 RequerAutenticacao: false,
	 },
 }