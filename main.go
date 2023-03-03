package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	//w.Write([]byte("hello"))
//	//_, _ = fmt.Fprintln(w, "<h1>hello go lang</h1><h2>how are you!</h2>")
//	b, err := ioutil.ReadFile("./hello.txt")
//	if err != nil {
//		fmt.Println("出大问题了")
//		return
//	}
//	_, _ = fmt.Fprintln(w, string(b))
//	//fmt.Println()
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":8088", nil)
//	if err != nil {
//		fmt.Printf("http serve failed, err:%v\n", err)
//		return
//	}
//}

func sayHello(c *gin.Context) {
	//c.HTML(200,"hello")
	c.JSON(200, gin.H{
		"message": "hello goalng",
	})
}

func main() {
	//r := gin.Default()
	//r.GET("/hello", sayHello)
	//r.GET("/book", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"method": "GET",
	//	})
	//})
	//r.POST("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "POST",
	//	})
	//})
	//r.PUT("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "PUT",
	//	})
	//})
	//r.DELETE("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "DELETE",
	//	})
	//})
	//r.Run(":9090")
	u := "http://accidentsrecordstest.wiltechs.com:8089"
	url := fmt.Sprintf("%s/api/accident/injuredPlatform/createTicket", u)
	fmt.Println(url)

	//_, err = conn.Do("HSet", "user_id","jwttoken","astoken")
}
