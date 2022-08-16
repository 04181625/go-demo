package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name"` //通常使用json
	Age  int    `json:"age" uri:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex"`
}

//对于输入的age绑定mustbig函数进行判断 小于等于18
func mustBig(fl validator.FieldLevel) bool {
	fmt.Println(fl.Field().Interface().(int))
	//通过反射的形式获取类型断言？在转成age的int形式 并进行判断验证规则是否通过
	if fl.Field().Interface().(int) <= 18 {
		return false
	}

	return true
}

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}

	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p) //body中json传递参数
		fmt.Println(err)
		//err := c.ShouldBindUri(&p) //http://localhost:8080/testBind/hkh/18/true
		//err := c.ShouldBindQuery(&p) //http://localhost:8080/testBind?name=hkh&age=22&sex=true
		if err != nil {
			c.JSON(200, gin.H{
				"msg":   "报错了",
				"error": err.Error(),
				"data":  gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功",
				"data": p,
			})
		}
	})

	r.Run()
}
