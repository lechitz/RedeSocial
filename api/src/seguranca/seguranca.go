package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash é a segurança da string para a senha
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha : compara uma senha e um hash e retorna se são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
