package security

import "golang.org/x/crypto/bcrypt"

func EncodePassword(password string) string {
	encodedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encodedPassword)
}

func IsPasswordValid(encodedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
}
