package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	//input := "04/17/2023 12:34"
	//regex := `^\d{2}/\d{2}/\d{4}\s\d{2}:\d{2}$`
	//if matched, err := regexp.MatchString(regex, input); err != nil {
	//	fmt.Println("正则表达式错误:", err)
	//} else if matched {
	//	fmt.Println("输入字符串符合 MM/DD/YYYY HH:mm 格式")
	//} else {
	//	fmt.Println("输入字符串不符合 MM/DD/YYYY HH:mm 格式")
	//}

	//inputString := "04/18/2023 00:31"                    // 输入的时间字符串
	//pst, err := time.LoadLocation("America/Los_Angeles") // PST（太平洋标准时间）时区
	//if err != nil {
	//	fmt.Println("加载时区信息失败:", err)
	//	return
	//}
	//inputTime, err := time.ParseInLocation("01/02/2006 15:04", inputString, pst)
	//if err != nil {
	//	inputTime, err = time.ParseInLocation("01/02/2006", inputString, pst)
	//	if err != nil {
	//		inputTime, err = time.ParseInLocation("15:04", inputString, pst)
	//		if err != nil {
	//			fmt.Println("%", inputString, "%")
	//			return
	//		}
	//	}
	//}
	//cst := time.FixedZone("CST", 8*60*60) // 中国标准时间时区
	//inputTime = inputTime.In(cst)
	//formattedTime := fmt.Sprintf("%%%s%%", inputTime.Format("01/02/2006 15:04"))
	//fmt.Println(formattedTime)
	inputString := "04/18/2023 00:4"                     // 输入的时间字符串
	pst, err := time.LoadLocation("America/Los_Angeles") // PST（太平洋标准时间）时区
	if err != nil {
		fmt.Println("加载时区信息失败:", err)
		return
	}
	pattern := `^(\d{2}/\d{2}/\d{4}\s\d{2}:\d{2})|(\d{2}/\d{2}/\d{4})|(\d{2}:\d{2})$`
	r := regexp.MustCompile(pattern)
	if r.MatchString(inputString) {
		inputTime, err := time.ParseInLocation("01/02/2006 15:04", inputString, pst)
		if err != nil {
			inputTime, err = time.ParseInLocation("01/02/2006", inputString, pst)
			if err != nil {
				inputTime, err = time.ParseInLocation("15:04", inputString, pst)
				if err != nil {
					fmt.Println("%", inputString, "%")
					return
				}
			}
		}
		cst := time.FixedZone("CST", 8*60*60) // 中国标准时间时区
		inputTime = inputTime.In(cst)
		formattedTime := fmt.Sprintf("%%%s%%", inputTime.Format("2006-01-02 15:04"))
		fmt.Println(formattedTime)
	} else {
		fmt.Println("%", inputString, "%")
	}
}
