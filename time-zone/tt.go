package main

import (
	"fmt"
	"time"
)

func main() {
	strTime := "2023-04-22 11:06:05 +0800"
	fmt.Println(convertTime(strTime))
}

func convertTime(strTime string) (string, error) {
	// 解析字符串時間
	t, err := time.Parse("2006-01-02 15:04:05 -0700", strTime)
	if err != nil {
		return "", err
	}

	// 轉換時區爲PST
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return "", err
	}
	t = t.In(loc)

	// 格式化爲目標格式的字符串時間
	return t.Format("15:04:05 01/02/2006"), nil
}
