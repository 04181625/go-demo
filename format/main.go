package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func main() {
//	t := time.Now()
//	ymd := t.Format("20060102") // 获取当前时间的年月日，格式为yyyymmdd
//	for i := 1; i <= 100; i++ { // 生成100个数字
//		num := strconv.Itoa(i)              // 将数字转换成字符串
//		numStr := fmt.Sprintf("%012s", num) // 将字符串补齐到12位
//		fmt.Println(ymd + numStr)           // 输出结果
//	}
//}

//func main() {
//	t := time.Now()
//	ymd := t.Format("20060102") // 获取当前时间的年月日，格式为yyyymmdd
//	for i := 1; i <= 100; i++ { // 生成100个数字
//		num := strconv.Itoa(i)              // 将数字转换成字符串
//		numLen := len(num)                  // 获取字符串的长度
//		zeroStr := "000000000000"           // 12个0的字符串
//		numStr := zeroStr[:12-numLen] + num // 将数字字符串前面补0，使其总长度为12位
//		fmt.Println(ymd + numStr)           // 输出结果
//	}
//}

func main() {
	//t := time.Now()
	//ymd := t.Format("20060102") // 获取当前时间的年月日，格式为yyyymmdd
	////for i := 1; i <= 99; i++ {  // 生成100个数字
	////	num := strconv.Itoa(i)             // 将数字转换成字符串
	////	numLen := len(num)                 // 获取字符串的长度
	////	zeroStr := "000"                   // 10个0的字符串
	////	numStr := zeroStr[:2-numLen] + num // 将数字字符串前面补0，使其总长度为10位
	////	fmt.Println(ymd + numStr)          // 输出结果
	////}
	//code := "105214"
	//badgeId := "2322"
	//var count uint = 4
	//countStr := strconv.Itoa(int(count)) //需要01-99
	//zero := "000"
	////a := zero[:2-len(countStr)]
	////mumStr := a + countStr
	//var builder1 strings.Builder
	//builder1.WriteString(ymd)
	//builder1.WriteString(code)
	//builder1.WriteString(badgeId)
	//builder1.WriteString(zero[:2-len(countStr)])
	//builder1.WriteString(countStr)
	//id := builder1.String()
	//fmt.Println(id)
	my := map[string]string{"1a": "Very", "2b": "good", "3c": "day"}

	a := mapToString(my)
	fmt.Println(a)

	var num int64 = 1234567890

	// 将 int64 类型转换为字符串类型
	str := strconv.FormatInt(num, 10)

	// 使用 strings.Split 函数对整数部分进行千分位分割
	parts := strings.Split(str, "")
	result := ""
	for i := len(parts) - 1; i >= 0; i-- {
		result = parts[i] + result
		if (len(parts)-i)%3 == 0 && i > 0 {
			result = "," + result
		}
	}

	fmt.Println(result)

}

func mapToString(m map[string]string) (a string) {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString(" like ")
		sb.WriteString(v)
		sb.WriteString(" and ")
	}
	a = sb.String()
	return a[:len(a)-4]
}
