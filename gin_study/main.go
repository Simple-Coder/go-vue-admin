package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.DefaultQuery("pwd", "123456")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "李四")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
