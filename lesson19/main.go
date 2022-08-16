package main

//测试数据库连接
import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(db)
	//defer db.
}
