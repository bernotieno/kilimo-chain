package asfuncss

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(m string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("ERROR")
		return ""
	}
	return string(hashedPassword)
}
