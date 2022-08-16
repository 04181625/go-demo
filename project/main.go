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

//确定数据库表模型model
//type Todo struct {
//	//gorm.Model
//	ID     int    `json:"id"`
//	Title  string `json:"title"`
//	Status bool   `json:"status"`
//}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//连接数据库
func initMySQL() (err error) {
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //有个全局的db这里用 =
	if err != nil {
		fmt.Println("error")
	}
	return err
}

//从post请求中获取参数name和password
//写入数据库？

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}) //模型和数据库表绑定 建表

	r := gin.Default()
	r.LoadHTMLFiles("project/login.html", "project/index.html")

	//login.html输入名字密码
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//r.GET("/update", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "update.html", nil)
	//})

	r.GET("/get", func(c *gin.Context) {
		//u1 := User{Name: "kkk", Password: "999"}
		//DB.Create(&u1)
		//u1.Name = "ojdoq"
		//DB.Create(&u1)
		var user User
		DB.First(&user)
		fmt.Println(user)
		user.Name = "cxk"
		DB.Debug().Save(&user)
	})

	//提交 获取请求参数
	r.POST("/login", func(c *gin.Context) {
		//index.html请求后获取文本蓝提交的参数
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
		fmt.Println(username)
		fmt.Println(password)
		var user User
		user.Name = username
		user.Password = password
		fmt.Println(user)
		//写入数据库
		if err = DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, user)
		}

	})

	r.PUT("/update", func(c *gin.Context) {
		//index.html请求后获取文本蓝提交的参数
		username := c.PostForm("username") //ab gg
		newname := c.PostForm("newname")   //12 ppp
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":    username,
			"Newname": newname,
		})
		fmt.Println(username) //ab gg
		fmt.Println(newname)  //12 ppp
		var user User
		//user.Name = newname

		if err = DB.Debug().Where("name=?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		//c.BindJSON(&user)
		fmt.Println("----")
		user.Name = newname
		fmt.Println(user)

		//DB.Debug().Model(&user).Update("name", newname)
		//DB.Debug().Save(&user)//save需要有where条件？
		//if err = DB.Debug().Save(&user).Error; err != nil {
		//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		//} else {
		//	c.JSON(http.StatusOK, user)
		//}

	})

	//获取userlist 查询数据库 select *
	r.GET("/userlist", func(c *gin.Context) {
		var userlist []User
		if err = DB.Find(&userlist).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, userlist)
		}

	})

	//r.PUT("/userupdate", func(c *gin.Context) {
	//	//DB.First(&user)
	//	//user.Name = "黄凯辉"
	//	//user.Number = "999"
	//	//DB.Debug().Save(&user)
	//	name, ok := c.Params.Get("name")
	//	if !ok {
	//		c.JSON(http.StatusOK, gin.H{"error": "name不存在"})
	//		return
	//	}
	//	newname, ok := c.Params.Get("newname")
	//	if !ok {
	//		c.JSON(http.StatusOK, gin.H{"error": "newname不存在"})
	//		return
	//	}
	//	fmt.Println(name)
	//	fmt.Println(newname)
	//	var user User
	//	if err = DB.Where("name=?", name).First(&user).Error; err != nil {
	//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//		return
	//	}
	//	c.BindJSON(&user)
	//	fmt.Println(user)
	//	if err = DB.Save(&user).Error; err != nil {
	//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//	} else {
	//		c.JSON(http.StatusOK, user)
	//	}
	//})

	//根据name进行delete操作
	r.DELETE("/user/:name", func(c *gin.Context) {
		name, ok := c.Params.Get("name")
		fmt.Println(name)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "name不存在"})
			return
		}
		if err = DB.Where("name=?", name).Delete(User{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{name: "deleted"})
		}
	})

	r.Run()
}
