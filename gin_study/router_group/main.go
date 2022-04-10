package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//洋葱模型
func middle1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是1")
		c.Next()
		fmt.Println("我在方法后，我是1")
	}
}
func middle2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是2")
		c.Next()
		fmt.Println("我在方法后，我是2")
	}
}

func main() {
	engine := gin.Default()
	v1 := engine.Group("v1").Use(middle1(), middle2())
	v1.GET("/test", func(c *gin.Context) {
		fmt.Println("我是分组方法的内部")
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	engine.Run(":8080")
}
