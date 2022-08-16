package main

//获取query参数
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//get请求url ？后是querystring参数
	//多个keyvalue形式用&连接参数
	r.GET("/web", func(c *gin.Context) {
		//获取浏览区请求携带的query string 参数
		name := c.Query("query")
		age := c.Query("age")
		//带默认值
		//name := c.DefaultQuery("query", "somebody")
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{ //通过query获取请求中携带的参数
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8080")
}
