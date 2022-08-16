package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//userinfo-->数据表

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	//root:4fabregas@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动创建表
	db.AutoMigrate(&UserInfo{})
	//defer db.Close

	//新建一条记录
	//u1 := UserInfo{1, "hkh", "男", "唱"}
	//u2 := UserInfo{2, "hkh", "男", "跳"}
	//u3 := UserInfo{3, "hkh", "男", "rap"}
	//db.Create(u3)

	//查询一条记录
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	//更新
	db.Model(&u).Update("Gender", "鸡你")
}
