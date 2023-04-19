package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "4/18/2023 10:1" // 输入的字符串
	parts := strings.Split(inputString, " ")
	if len(parts) != 2 {
		fmt.Println("输入字符串格式不正确")
		return
	}
	date := parts[0]    // 日期部分
	timeStr := parts[1] // 时间部分
	dateParts := strings.Split(date, "/")
	timeParts := strings.Split(timeStr, ":")
	if len(timeParts) != 2 {
		timeParts = append(timeParts, "00")
	}
	for i := range dateParts {
		dateParts[i] = fmt.Sprintf("%02s", dateParts[i])
	}
	for i := range timeParts {
		timeParts[i] = fmt.Sprintf("%02s", timeParts[i])
	}
	formattedString := fmt.Sprintf("%s/%s/%s %s:%s", dateParts[0], dateParts[1], dateParts[2], timeParts[0], timeParts[1])
	fmt.Println(formattedString)
}
