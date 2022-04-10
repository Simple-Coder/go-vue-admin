package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.JSON(200, gin.H{
			"msg": file,
		})
	})

	engine.Run(":8080")
}
