package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	//dsn := "root:4fabregas@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//gorm.Open("mysql", dsn)
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	return err

}

func main() {
	//创建数据库
	//连接
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//需要告诉gin去哪里找静态文件css js
	r.Static("/static", "./bubble/static")
	//需要告诉gin去哪里找indexhtml
	r.LoadHTMLGlob("bubble/template/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//待办事项 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端洁面填写代办事项 点击提交 发送到这里
			//从请求中拿出数据
			//存入数据库
			//返回响应
			var todo Todo
			c.BindJSON(&todo)
			fmt.Println(todo)
			//err := DB.Create(&todo).Error
			//if err != nil{
			//	c.JSON(http.StatusOK,gin.H{"error": })
			//}
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}

		})
		//查看一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个代办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			fmt.Println(todo)
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	r.Run()

}
