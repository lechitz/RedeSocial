package modelos

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

//Modelo de usuário na RedeSocial
type Usuario struct {
	ID uint64 `json:"id,omitempty"` //omitempty não deixa passar o ID em branco para o JSON
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

//Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	//Verificando se os campos foram preenchidos
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	return nil
}

//validar verifica se todos os campos da struct estão preenchidos (exceto o ID e o CriadoEm)
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome de usuário é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("O nick de usuário é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("O email de usuário é obrigatório")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha de usuário é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	//Retirando os espaços das extremidades
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}