package main

import (
	"fmt"
	"time"
)

func main() {
	str1 := "4/18/2023 10:1"
	t1, err1 := time.Parse("1/2/2006 15:4", str1)
	if err1 != nil {
		fmt.Println("解析错误：", err1)
		return
	}
	if t1.Minute() < 10 {
		t1 = t1.Add(time.Minute * time.Duration(10-t1.Minute()))
	}
	if t1.Hour() < 10 {
		t1 = t1.Add(time.Hour * time.Duration(10-t1.Hour()))
	}
	result1 := t1.Format("1/2/2006 15:04")
	fmt.Println("输入的字符串：", str1)
	fmt.Println("填充后的字符串：", result1)
}
