package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//bind反射结构体
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {

	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		var u UserInfo //声明一个userinfo类型的变量u 需要传递结构体指针 否则修改读取的是值的副本
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

		//fmt.Printf("%#v\n", u)
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//	//"username": username,
		//	//"password": password,
		//})
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo //声明一个userinfo类型的变量u
		//需要传递结构体指针 否则修改读取的是值的副本
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}
