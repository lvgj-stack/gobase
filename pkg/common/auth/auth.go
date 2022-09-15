package auth

import "golang.org/x/crypto/bcrypt"

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}