package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//方法1使用map
		//data := map[string]interface{}{
		//	"name":    "小肥猪",
		//	"message": "hellow orld",
		//	"age":     18,
		//}
		data := gin.H{"name": "飞猪", "message": "helloworld", "age": 19}
		c.JSON(http.StatusOK, data)
	})

	type msg struct {
		//如果小写开头就包外用不了
		//反斜杠
		Name    string `json:"name"`
		Message string `bson:"message"`
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "汉堡",
			Message: "helasdjl",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
