package main

import (
	"fmt"
	"time"
)

func main() {
	//loc := time.Now().Location()
	//fmt.Println(loc)
	//fmt.Println(time.Now())
	//
	//now := time.Now()
	//fmt.Println(now)
	//fmt.Println("===")
	//
	//loc1, _ := time.LoadLocation("UTC")
	//fmt.Println(now.In(loc1))
	//
	//loc2, _ := time.LoadLocation("Europe/Berlin")
	//fmt.Println(now.In(loc2)) // 2022-07-18 15:19:59.9636 +0200 CEST
	//
	//loc3, _ := time.LoadLocation("America/New_York")
	//fmt.Println(now.In(loc3)) // 2022-07-18 09:19:59.9636 -0400 EDT
	//
	//loc4, _ := time.LoadLocation("Asia/Dubai")
	//fmt.Println(now.In(loc4)) // 2022-07-18 17:19:59.9636 +0400 +04
	//
	//loc5, _ := time.LoadLocation("America/Los_Angeles")
	//fmt.Println(now.In(loc5)) // 2022-07-18 17:19:59.9636 +0400 +04
	//
	//a := "2022-06-18 17:19:59.9636 UTC"
	//layout := "2006-01-02 15:04:05 MST"
	//t, err := time.Parse(layout, a)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(t)
	//
	//loc6, _ := time.LoadLocation("America/New_York")
	//fmt.Println(t.In(loc6)) // 2022-07-18 17:19:59.9636 +0400 +04

	//timeZone := "America/New_York"
	//now := time.Now()
	//loc, _ := time.LoadLocation(timeZone)
	//fmt.Println(now.In(loc))
	//now := time.Now()
	//start, _ := time.Parse("01/02/2006", now.Format("2006-01-02"))
	//fmt.Println(now)
	//fmt.Println(start)
	//a := 124
	//b := 4563
	//d := 100
	//c := uint64(a) + uint64(b)
	//fmt.Println(c)
	//fmt.Println(reflect.TypeOf(c))
	//fmt.Println(uint64(a) + uint64(d))
	//str := "This is a Y1P test string with Upper and Lowercase letters"
	//
	//contains := strings.Contains(str, "Y1P") || strings.Contains(str, "SJ88") ||
	//	strings.Contains(str, "FTR") || strings.Contains(str, "VM") ||
	//	strings.Contains(str, "OMG") || strings.Contains(str, "AUD")
	//
	//hasUpper := false
	//hasLower := false
	//
	//for _, ch := range str {
	//	if unicode.IsUpper(ch) {
	//		hasUpper = true
	//	}
	//	if unicode.IsLower(ch) {
	//		hasLower = true
	//	}
	//}
	//
	//fmt.Println(contains) // 输出true
	//fmt.Println(hasUpper) // 输出true
	//fmt.Println(hasLower) // 输出true
	//date1 := "2022-01-01 00:00:00"
	//date2 := "2022-01-10 00:00:00"

	// 解析日期时间字符串为时间对象
	//t1, err := time.Parse("2006-01-02 15:04:05", date1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//t2, err := time.Parse("2006-01-02 15:04:05", date2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// 计算两个时间之间的天数差异
	//days := int(t2.Sub(t1).Hours() / 24)
	//
	//fmt.Printf("天数差异：%d", days)
	//
	//fmt.Println(time.Now())
	//t := time.Now()

	// 将时间格式化为RFC3339格式的字符串
	//tFormatted := t.Format(time.RFC3339)

	//t := time.Now()
	//
	//// 将时间格式化为RFC3339格式的字符串
	//tFormatted := t.Format(time.RFC3339)
	//
	//// 将时间字符串解析为时间对象
	//tRFC3339, err := time.Parse(time.RFC3339, tFormatted)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(tRFC3339)
	//t2 := time.Now()
	//t2, err := time.Parse(time.RFC3339, time.Now().String())
	//if err != nil {
	//	fmt.Println("转化失败：", err)
	//}
	//fmt.Println(t2)
	//strTime := "2023-11-24 06:30:00"
	//
	//// 时间格式
	//timeLayout := "2006-01-02 15:04:05"
	//
	//// 转换为time类型
	//t, err := time.Parse(timeLayout, strTime)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//est, _ := time.LoadLocation("America/New_York")
	//cst, _ := time.LoadLocation("America/Chicago")
	//mst, _ := time.LoadLocation("America/Denver")
	//pst, _ := time.LoadLocation("America/Los_Angeles")
	//akst, _ := time.LoadLocation("America/Anchorage")
	//hst, _ := time.LoadLocation("Pacific/Honolulu")
	//now := time.Now()
	//// 打印每个时区的时间
	//fmt.Println(t)
	//fmt.Println("Eastern Time (EST/EDT):", t.In(est).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Central Time (CST/CDT):", t.In(cst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Mountain Time (MST/MDT):", t.In(mst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Pacific Time (PST/PDT):", t.In(pst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Alaska Time (AKST/AKDT):", t.In(akst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Hawaii Time (HST/HDT):", t.In(hst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("=============================================================================================")
	//fmt.Println(now)
	//fmt.Println("Eastern Time (EST/EDT):", now.In(est).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Central Time (CST/CDT):", now.In(cst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Mountain Time (MST/MDT):", now.In(mst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Pacific Time (PST/PDT):", now.In(pst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Alaska Time (AKST/AKDT):", now.In(akst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("Hawaii Time (HST/HDT):", now.In(hst).Format("2006-01-02 15:04:05 MST"))
	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//
	//for i := 0; i < 10; i++ {
	//	var input string
	//	fmt.Println("输入:")
	//	fmt.Scan(&input)
	//	output := Transform(input)
	//	fmt.Println(output)
	//}

	//str := "2023-04-04T13:35:48+08:00"
	//t, err := time.Parse(time.RFC3339, str)
	//if err != nil {
	//	panic(err)
	//}
	//localTime := t.In(time.Local)
	//fmt.Println(localTime)
	//fmt.Println(localTime.Format("01/02/2006"))
	//str := "2023-04-04 13:56:09 +0800"
	//t, err := time.Parse("2006-01-02 15:04:05 -0700", str)
	//if err != nil {
	//	panic(err)
	//}
	//localTime := t.In(time.Now().Location())
	//fmt.Println(localTime)
	str := "04/03/2022"
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("01/02/2006", str, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}

func Transform(zone string) string {
	est, _ := time.LoadLocation("America/New_York")
	cst, _ := time.LoadLocation("America/Chicago")
	mst, _ := time.LoadLocation("America/Denver")
	pst, _ := time.LoadLocation("America/Los_Angeles")
	akst, _ := time.LoadLocation("America/Anchorage")
	hst, _ := time.LoadLocation("Pacific/Honolulu")
	now := time.Now()
	switch zone {
	case "1":
		return now.In(est).Format("2006-01-02 15:04:05 MST")
	case "2":
		return now.In(cst).Format("2006-01-02 15:04:05 MST")
	case "3":
		return now.In(mst).Format("2006-01-02 15:04:05 MST")
	case "4":
		return now.In(pst).Format("2006-01-02 15:04:05 MST")
	case "5":
		return now.In(akst).Format("2006-01-02 15:04:05 MST")
	case "6":
		return now.In(hst).Format("2006-01-02 15:04:05 MST")
	}

	return ""
}
