package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard

	Teachers []Teacher `gorm:"many2many:student_teachers;"` //多对多关系
	//TeacherID uint
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	//StudentID uint
	Students []Student `gorm:"many2many:student_teachers;"`
}

func main() {
	dsn := "root:4fabregas@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})
	i := IDCard{
		Num: 123456,
	}
	t := Teacher{
		TeacherName: "老师傅",
		//Students:    []Student{s},
	}
	s := Student{
		StudentName: "qm",
		IDCard:      i,
		Teachers:    []Teacher{t},
	}
	//i := IDCard{ //按照结构体内顺序创建 否则报错？
	//	Num: 123,
	//}
	//t := Teacher{
	//	TeacherName: "老师傅",
	//	//Students:    []Student{s},
	//}
	c := Class{
		ClassName: "计算机的班级",
		Students:  []Student{s},
	}

	_ = db.Debug().Create(&c).Error
	_ = db.Debug().Create(&t).Error

}
