package security

import "golang.org/x/crypto/bcrypt"

// HASH x BCRYPT = BCRYPT AINDA PODE SER DESCRIPTOGRAFADO.. HASH É INRREVERSÍVEL

// Existe uma função que compara uma string com um bash e verificar se tem o mesmo valor

// function Hash received string and set a hash for it
func Hash(password string) ([]byte, error) {
	// second parameter is used to increase security
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// compare password and password with hash and return if for equals
func VerifyPassword(password, passwordWithHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(password))
}
