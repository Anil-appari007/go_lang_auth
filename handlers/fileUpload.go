package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	folder := "Downloads"
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		destFile := folder + "/" + file.Filename
		fmt.Println(file.Filename)
		fmt.Printf("Saving file to %s", destFile)
		c.SaveUploadedFile(file, destFile)
	}
	// c.String(http.statusOk, "file %s uploaded successfully")
	c.JSON(200, gin.H{"success": "file uploaded"})
}
