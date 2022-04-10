package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//engine.POST("/testUpload", func(c *gin.Context) {
	//	file, _ := c.FormFile("file")
	//	c.SaveUploadedFile(file, "./"+file.Filename)
	//	c.JSON(200, gin.H{
	//		"msg": file,
	//	})
	//})
	//

	//文件写回
	engine.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.Writer.
			Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
		c.File("./" + file.Filename)
	})

	engine.Run(":8080")
}
