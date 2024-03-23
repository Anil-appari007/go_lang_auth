package utils

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ExitIfErr(err error, method string) {
	if err != nil {
		fmt.Println("\nError During %s", method)
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
func CheckError(err error, method string) {
	if err != nil {
		fmt.Printf("\nError During %s\n", method)
		fmt.Println(err)
		return
		// os.Exit(2)
	}
}

func ReturnIfError(err error, status int, c *gin.Context, msg string) bool {
	if err != nil {
		// message := map[string]string{
		// 	"error": string(err.Error()),
		// }
		fulllError := msg + string(err.Error())
		// message := map[string]string{
		// 	"error": msg,
		// }
		message := map[string]string{
			"error": fulllError,
		}
		c.JSON(status, message)
		return true
	} else {
		return false
	}

}
