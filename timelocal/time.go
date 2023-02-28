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
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建一个表示当前时间的 time.Time 对象
	//t := time.Now()
	a := "2023-01-28 03:48:30"
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, a)
	// 将该时间转换为 "America/New_York" 时区的时间
	tInLoc := t.In(loc)
	// 检查该时间是否处于夏令时
	isDST := tInLoc.IsDST()

	fmt.Printf("当前时间 %s 是否处于夏令时：%v\n", tInLoc, isDST)
	// 判断是否处于夏令时
	//isDST := location.Tzname(newYorkTime) != location.Tzname(localTime)
	//if isDST {
	//	fmt.Println("It is currently Daylight Saving Time in New York.")
	//} else {
	//	fmt.Println("It is currently Standard Time in New York.")
	//}
}
