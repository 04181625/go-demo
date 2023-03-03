package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前本地时间
	//localTime := time.Now()
	//
	//// 加载指定时区的时区规则
	//location, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//	fmt.Println("Error loading location:", err)
	//	return
	//}
	//
	//// 将本地时间转换为指定时区的时间
	//newYorkTime := localTime.In(location)
	//
	//// 打印转换后的时间
	//fmt.Println("New York Time:", newYorkTime.Format(time.RFC3339))
	//loc, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// 创建一个表示当前时间的 time.Time 对象
	//t := time.Now()
	//a := "2023-01-28 03:48:30"
	//layout := "2006-01-02 15:04:05"
	//t, _ := time.Parse(layout, a)
	// 将该时间转换为 "America/New_York" 时区的时间
	//tInLoc := t.In(loc)
	// 检查该时间是否处于夏令时
	//isDST := tInLoc.IsDST()
	//fmt.Printf("当前时间 %s 是否处于夏令时：%v\n", tInLoc, isDST)
	// 判断是否处于夏令时
	//isDST := location.Tzname(newYorkTime) != location.Tzname(localTime)
	//if isDST {
	//	fmt.Println("It is currently Daylight Saving Time in New York.")
	//} else {
	//	fmt.Println("It is currently Standard Time in New York.")
	//}
	now := time.Now() //本地时区时间 pc右下角
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Day())
	timestamp1 := now.Unix()
	fmt.Println(timestamp1)
	t := time.Unix(1677822553, 0)
	fmt.Println(t)

	thisdate := "2014-03-17 14:55:06"
	timeformatdate, _ := time.Parse(datetime, thisdate)
	fmt.Println(timeformatdate)
	convdate := timeformatdate.Format(date)
	convshortdate := timeformatdate.Format(shortdate)
	convtime := timeformatdate.Format(times)
	convshorttime := timeformatdate.Format(shorttime)
	convnewdatetime := timeformatdate.Format(newdatetime)
	convnewtime := timeformatdate.Format(newtime)
	fmt.Println(convdate)
	fmt.Println(convshortdate)
	fmt.Println(convtime)
	fmt.Println(convshorttime)
	fmt.Println(convnewdatetime)
	fmt.Println(convnewtime)

	//GMT格林威治时
	//UTC协调世界时
	//
}

const (
	date        = "2006-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:02"
	shorttime   = "15:04"
	datetime    = "2006-01-02 15:04:05"
	newdatetime = "2006/01/02 15~04~02"
	newtime     = "15~04~02"
)
