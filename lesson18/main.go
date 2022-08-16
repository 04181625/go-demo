package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//中间件函数

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

//c.next调用后续的处理函数
//c.abort阻止调用后续函数

func m2(c *gin.Context) {
	fmt.Println("m2 in..")
	c.Next()
	fmt.Println("m2 out...")
}

func main() {
	r := gin.Default()
	//r.Use(m1) //全剧注册中间件
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg":"index",
	//	})
	//})
	r.Use(m1, m2)
	//r.GET("/index", m1, indexHandler)
	r.GET("/index", indexHandler)
	//运行顺序
	//先m1 m1-in c.next m2-in c.next indexhandle 执行完返回m2的next 然后是m2-out 执行完返回m1的next 然后是m1-out
	//m2 c.abort不执行后续的indexhandle 继续m2out

	r.Run()
}
