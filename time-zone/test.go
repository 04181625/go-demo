package main

import (
	"fmt"
	"strings"
	"time"
)

// 现在需要把输入的04/15/2023 17:54格式的字符串时间转换为2023-04-15T17:54:25-07:00格式的pst时间
// 把2023-04-15T17:54:25-07:00格式的pst时间转换为2006-01-02 15:04格式的cst时间
// GetTime后作为search数据库的条件
func main() {
	//fmt.Println(GetTime("2023-04-16T22:57:25-07:00")) //pst时间
	a := Tr("04/17/2023 22:40")
	fmt.Println(a)
	//fmt.Println(GetStringTimeFormMMDDYYYYRFC3339("2023-04-16T23:49:05-07:00"))
}

// Tr 04/15/2023 17:54 -> 2023-04-15T10:54:00-07:00 相当于utc时间变成了pst时间 -7h
func Tr(date string) string {
	t, err := time.Parse("01/02/2006 15:04", date)
	if err != nil {
		panic(err)
	}
	t1 := t.Format(time.RFC3339)
	t2 := t1[0:19]
	var sb strings.Builder
	sb.WriteString(t2)
	sb.WriteString("-07:00")
	return GetTime(sb.String())
}

// GetTime 2023-04-16T19:40:25-07:00 -> 2023-04-17 13:40 变成中国时间了
func GetTime(date string) string {
	if date != "" {
		t, err := time.Parse(time.RFC3339, date) //UTC时间
		if err != nil {
			panic(err)
		}
		bj, _ := time.LoadLocation("Asia/Shanghai")
		timeForm := t.In(bj)
		return timeForm.Format("2006-01-02 15:04")
	}
	return ""
}

func GetStringTimeFormMMDDYYYYRFC3339(date string) string { //CST
	inputFormat := "2006-01-02T15:04:05-07:00"
	if date != "" {
		t, err := time.Parse(inputFormat, date) //UTC时间
		if err != nil {
			panic(err)
		}
		return t.Format("01/02/2006")
	}
	return ""
}
