package main

import (
	"fmt"
	"time"
)

func main() {
	//password := "9999"
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//if err != nil {
	//	fmt.Println("生成哈希密码失败:", err)
	//	return
	//}
	//fmt.Println(hashedPassword)
	//fmt.Println(string(hashedPassword))
	//fmt.Println("helloworld!!!")
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
	fmt.Println(GetTime("2023-04-15T00:11:25-07:00")) //pst时间
	//fmt.Println(GetTimeFormMMDDYYYY("01/04/2009"))
	//fmt.Println(GetTime("2009-01-04 00:00:00 -0800"))

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
func GetTimeFormMMDDYYYYHHmm(date string) string { //CST
	if date != "" {
		t, err := time.Parse(time.RFC3339, date) //UTC时间
		if err != nil {
			panic(err)
		}
		timeForm := t.In(time.Now().Location())
		return timeForm.Format("01/02/2006 15:04")
	}
	return ""
}

func GetTimeFormMMDDYYYYLocalTime(date string) string { //CST
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", date) //UTC时间
	if err != nil {
		panic(err)
	}
	timeForm := t.In(time.Now().Location())
	return timeForm.Format("01/02/2006")
}

func GetTimeFormMMDDYYYY(date string) string {
	t, _ := time.ParseInLocation("01/02/2006", date, time.Local)
	timeForm := t.Format("2006-01-02")
	return timeForm
}
