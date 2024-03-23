package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string, jwtSecret []byte, tokenValidity int) (string, error) {
	// fmt.Println("Creating Token")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Duration(tokenValidity) * time.Second).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
