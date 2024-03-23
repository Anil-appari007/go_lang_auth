package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(Password string) []byte {
	var pwd = []byte(Password)
	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, 10)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(hashedPwd)
	return hashedPwd
}
