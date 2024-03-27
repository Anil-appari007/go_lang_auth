package handlers

import (
	"fmt"
	utils "go_lang_auth/Utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("mySecretToken")

func ValidateToken(c *gin.Context) {
	AuthToken := c.Request.Header.Get("Authorization")

	if AuthToken == "" {
		fmt.Println("AuthToken is empty")
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
	fmt.Println(AuthToken)
	splitToken := strings.Split(AuthToken, " ")
	fmt.Println(splitToken)
	Token := splitToken[1]
	fmt.Println(Token)
	err := utils.ParseToken(Token, jwtSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalidToken"})
		return
	}
	c.Next()
}
