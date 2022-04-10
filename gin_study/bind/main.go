package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//test bind

type PostParams struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required,mustMoreThan"`
	//Age int `json:"age" form:"age"`
}

func mustMoreThan(f1 validator.FieldLevel) bool {
	if f1.Field().Int() <= 18 {
		return false
	}
	//if (f1.Field().Interface().(int) < = 18) {
	return true
}
func main() {
	engine := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustMoreThan", mustMoreThan)
	}

	engine.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		//err := c.ShouldBindQuery(&p)
		if err != nil {
			c.JSON(500, gin.H{
				"msg":  "报错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了！",
				"data": p,
			})
		}
	})

	engine.Run(":8080")
}
