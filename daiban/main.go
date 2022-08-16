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

//Toso Model
type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/daiban?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	return err
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Todo{}) //模型和数据库表绑定 建表

	r := gin.Default()

	r.POST("/addtodo", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)
		//写入数据库
		if err = DB.Create(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}

	})

	//查询所有事项返回
	r.GET("/todolist", func(c *gin.Context) {
		var todolist []Todo
		if err = DB.Find(&todolist).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todolist)
		}

	})

	r.DELETE("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		fmt.Println(id)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
			return
		}
		if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{id: "was deleted"})
		}
	})

	r.Run()
}
