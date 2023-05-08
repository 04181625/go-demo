package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	password := "9999" //$2a$10$qJZmW70FKLMorB6mt0iN.efr5ibX9FvQnxjQkdaM/.Jm8sxbTpBHm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("生成哈希密码失败:", err)
		return
	}
	fmt.Println(hashedPassword)
	fmt.Println(string(hashedPassword))
	fmt.Println("helloworld!!!")
	//
	//password = "8888"
	//err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	//if err != nil {
	//	fmt.Println("error!!!!!!!")
	//	return
	//}
	//authorized := false
	//item := "/api/app/roles/users"
	//path := "/api/app/roles/users/124h12b239图片等nasa"
	//if strings.HasPrefix(path, item) {
	//	authorized = true
	//	fmt.Println(authorized)
	//	fmt.Println("hello true")
	//} else {
	//	fmt.Println(authorized)
	//	fmt.Println("hello false")
	//}
	//inputTime := "2016-07-01T00:00:00+08:00"
	//inputFormat := "2006-01-02T15:04:05-07:00"
	//t, err := time.Parse(inputFormat, inputTime)
	//if err != nil {
	//	fmt.Println("解析时间失败:", err)
	//	return
	//}
	//outputFormat := "01/02/2006"
	//outputTime := t.Format(outputFormat)
	//fmt.Println("输出时间字符串:", outputTime)
	// 输入的时间字符串

	//inputTime := "03/27/2023 00:06"
	//inputFormat := "01/02/2006 15:04"
	//t, err := time.Parse(inputFormat, inputTime)
	//if err != nil {
	//	fmt.Println("解析时间失败:", err)
	//	return
	//}
	//pst := time.FixedZone("PST", -8*60*60)
	//pstTime := t.In(pst)
	//cst := time.FixedZone("CST", -6*60*60)
	//cstTime := pstTime.In(cst)
	//outputFormat := "01/02/2006 15:04"
	//outputTime := cstTime.Format(outputFormat)
	//fmt.Println("输出时间字符串:", outputTime)

	//fmt.Println(GetTimeFormMMDDYYYYHHmm("2023-03-27T15:06:25+08:00"))
	//fmt.Println(GetTimeFormMMDDYYYYHHmm("2023-04-15T14:38:25+08:00"))
	//fmt.Println(GetTimeFormMMDDYYYY("01/04/2009"))
	//fmt.Println(GetTime("2009-01-04 00:00:00 -0800"))
	//fmt.Println(TramsFor("04/15/2023 17:53"))
	//fmt.Println(GetTime("2023-04-15 17:54:00 -0700 PDT"))
	//fmt.Println(GetTime("2023-04-16T22:40:25-07:00")) //pst时间

	// 定义输入字符串和输入格式
	//input := "04/16/2023 19:40:25"
	//inputFormat := "01/02/2006 15:04:05"
	//t, err := time.Parse(inputFormat, input) //2023-04-16 19:40:25 +0000
	//if err != nil {
	//	fmt.Println("解析时间错误:", err)
	//	return
	//}
	//location, err := time.LoadLocation("America/Los_Angeles")
	//if err != nil {
	//	fmt.Println("加载时区错误:", err)
	//	return
	//}
	//t = t.In(location) //2023-04-16 12:40:25 -0700
	//outputFormat := "2006-01-02T15:04:05-07:00"
	//output := t.Format(outputFormat) //2023-04-16T12:40:25-07:00
	//fmt.Println("输出时间:", output)
}

func Tr(date string) string {
	t, err := time.Parse("01/02/2006 15:04", date)
	if err != nil {
		panic(err)
	}
	local := time.Now().Location()
	nowInLocal := t.In(local)
	fmt.Println(nowInLocal.Format(time.RFC3339))
	return ""
}

//现在需要把输入的04/15/2023 17:54格式的字符串时间转换为2023-04-15T17:54:25-07:00格式的pst时间
// 把2023-04-15T17:54:25-07:00格式的pst时间转换为2006-01-02 15:04格式的cst时间
// GetTime后作为search数据库的条件

func TramsFor(date string) string {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("01/02/2006 15:04", date, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	return ""
}

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

func T() string {
	input := "04/16/2023 19:40:25"
	inputFormat := "01/02/2006 15:04:05"
	t, err := time.Parse(inputFormat, input) //2023-04-16 19:40:25 +0000
	if err != nil {
		fmt.Println("解析时间错误:", err)
		return ""
	}
	outputFormat := "2006-01-02T15:04:05-07:00"
	output := t.Format(outputFormat) //2023-04-16T12:40:25-07:00
	fmt.Println(output)
	fmt.Println(GetTime(output))
	return ""
}
