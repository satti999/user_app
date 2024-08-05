package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
func CheckHash(hpass string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hpass), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
