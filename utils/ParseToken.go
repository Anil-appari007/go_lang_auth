package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string, jwtSecret []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		// fmt.Println("Checking for Sign method")
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		// fmt.Println(err.Error())
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token is not valid")
	}
	return nil
}
