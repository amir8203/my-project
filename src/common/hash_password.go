package common

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(hashedPassword), err
}