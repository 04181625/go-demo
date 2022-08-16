package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Number string
}

type Haha struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Haha{})
	//u1 := User{Name: "hkh", Age: 1, Number: "123"}
	//DB.Create(&u1)
	//u2 := User{Name: "qwx", Age: 2, Number: "456"}
	//DB.Create(&u2)
	//u3 := User{Name: "cy", Age: 3, Number: "789"}
	//DB.Create(&u3)
	//u4 := User{Name: "lyf", Age: 4, Number: "100"}
	//DB.Create(&u4)
	//查询
	//var user User
	//DB.First(&user)//第一条
	//查询表中所有
	//var users []User
	//DB.Find(&users)
	//fmt.Printf("user:%#v\n", users)

	a1 := Haha{Name: "qwe", Password: "99"}
	DB.Create(&a1)
	var haha Haha
	DB.Debug().First(&haha)
	fmt.Println(haha)
	haha.Name = "jkl"
	DB.Debug().Save(&haha)
	//DB.First(&user)
	//user.Name = "jkl"
	//user.Number = "999"
	//DB.Debug().Save(&user)

	//DB.Model(&User{}).Where("name = ?", "jkl").Update("name", "op")

	//DB.Debug().Model(&user).Update("name", "wqouheibq")
	//DB.Save(&user)

}
