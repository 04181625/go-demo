package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//http://localhost:8080/path/123?user=qm&pwd=123456
	//path后面的id=123用param获取，user和pwd是地址栏后面的通过query传参的
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "hkh")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			//"success": true,
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})

	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "hkh")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			//"success": true,
			"user": user,
			"pwd":  pwd,
		})
	})

	r.DELETE("/path/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.PUT("/path/", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "hkh")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			//"success": true,
			"user": user,
			"pwd":  pwd,
		})
	})

	r.Run()

}
