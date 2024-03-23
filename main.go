package main

import (
	"database/sql"
	"fmt"
	utils "go_lang_auth/Utils"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("mySecretToken")
var tokenValidity = 300 // seconds

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var newuser credentials
	if err := c.BindJSON(&newuser); err != nil || newuser.Username == "" || newuser.Password == "" {
		c.JSON(400, gin.H{"error": "invalid json data"})
		return
	}

	// Check if user exists
	db := utils.OpenDbConn()
	rows, err := db.Query("SELECT username from users where username = $1", newuser.Username)
	utils.CheckError(err, "user check")
	if rows.Next() {
		fmt.Printf("User %s already exists\n", newuser.Username)
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	} else {
		fmt.Println("\nadding User ")
	}

	// encrypt password
	hashedPwd := utils.EncryptPwd(string(newuser.Password))
	fmt.Println(string(hashedPwd))
	_, err = db.Query("INSERT INTO users values($1, $2)", newuser.Username, hashedPwd)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "unable to add user"})
		return
	}
	c.JSON(201, gin.H{"success": "user added"})

}
func LoginUser(c *gin.Context) {
	var creds credentials
	if err := c.BindJSON(&creds); err != nil || creds.Username == "" || creds.Password == "" {
		c.JSON(400, gin.H{"error": "invalid json data"})
		return
	}
	db := utils.OpenDbConn()

	rows := db.QueryRow("select password from users where username = $1;", string(creds.Username))

	var hashedPwd string

	err := rows.Scan(&hashedPwd)
	if err != nil && err == sql.ErrNoRows {
		if err == sql.ErrNoRows {
			utils.ReturnIfError(err, 400, c, "user does not exist")
			return
		}

	}

	// fmt.Println("Encrypted Pwd %s", hashedPwd)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(creds.Password))
	if utils.ReturnIfError(err, 400, c, "Incorrect Password") {
		return
	}
	// Create JWT Token and send
	token, err := utils.CreateToken(creds.Username, jwtSecret, tokenValidity)
	if utils.ReturnIfError(err, 400, c, "during token creation") {
		return
	}
	userData := map[string]string{
		"success": "user logged",
		"token":   token,
	}

	// c.JSON(200, gin.H{"success": "user logged in"})
	c.JSON(200, userData)
}

func GetItems(c *gin.Context) {
	AuthToken := c.Request.Header.Get("Authorization")
	// fmt.Println(AuthToken)
	splitToken := strings.Split(AuthToken, " ")
	// fmt.Println(splitToken)
	Token := splitToken[1]
	// fmt.Println(Token)
	err := utils.ParseToken(Token, jwtSecret)
	if utils.ReturnIfError(err, 400, c, "Auth Token Parse") {
		return
	}
	c.JSON(200, gin.H{"data": "testdata"})
}
func checkDb() {
	db := utils.OpenDbConn()
	row := db.QueryRow("SELECT version();")
	var version string
	err := row.Scan(&version)
	utils.ExitIfErr(err, "version scan")
	// fmt.Println(version)
	fmt.Println("DB is connected")
}
func main() {
	checkDb()
	router := gin.Default()
	router.POST("/signup", CreateUser)
	router.POST("/signin", LoginUser)

	router.GET("/items", GetItems)
	router.Run("localhost:8082")
}
