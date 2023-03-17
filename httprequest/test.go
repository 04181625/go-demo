package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type User struct {
//	ID       int    `json:"id"`
//	Name     string `json:"name"`
//	Username string `json:"username"`
//	Email    string `json:"email"`
//}

type TTest struct {
	Count string `json:"count"`
}

func main() {
	// 构造 HTTP GET 请求
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("HTTP GET 请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取 HTTP 响应体并解析为 JSON 格式
	//var users []User
	var test []TTest
	//var todos []map[string]interface{}
	//err = json.NewDecoder(resp.Body).Decode(&todos)
	err = json.NewDecoder(resp.Body).Decode(&test)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	//fmt.Println(len(todos))
	fmt.Println(len(test))

	// 输出结果
	//for _, user := range users {
	//	fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	//}
}
