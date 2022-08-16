package main

//获取表单提交的参数
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取form参数
func main() {

	r := gin.Default()
	r.LoadHTMLFiles("lesson12/login.html", "lesson12/index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//login post
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
		fmt.Println(username)
		fmt.Println(password)
	})
	r.Run()
}
