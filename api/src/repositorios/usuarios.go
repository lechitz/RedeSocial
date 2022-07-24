package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios Recebe um banco aberto pelo controller(ele que chama essa função). Joga essa função dentro do struct de usuário (basicamente criando uma instância com o banco que foi aberto)
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64,error){
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?,?,?,?)",
		)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	//resultado insere um usuário no banco
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//Buscar traz todos usuários pelo nomeOuNick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %%nomeOuNick%%

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm from usuarios where nome  LIKE ? or nick LIKE ?",nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			); erro != nil {
				return nil, erro
			}

			usuarios = append(usuarios,usuario)
	}

	return usuarios, nil
}

//BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		//return com modelos.Usuario{} por conta do retorno esperado, passando {} vazio
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	//Criando o usuário
	var usuario modelos.Usuario

	//Se tiver alguma linha para ser lida
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Atualizar modifica o usuario no banco
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome=?, nick=?, email=? where id=?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
	
}

//Deletar exclui as informações de um usuario no banco
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}