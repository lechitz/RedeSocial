package modelos

import "time"

//Modelo de usuário na RedeSocial
type Usuario struct {
	ID uint64 `json:"id,omitempty"` //omitempty não deixa passar o ID em branco para o JSON
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}