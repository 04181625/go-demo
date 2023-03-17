package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//	type Car struct {
//		CreateDate          string `json:"CreateDate"`
//		VehicleSerialNumber string `json:"VehicleSerialNumber"`
//		Year                string `json:"Year"`
//		Make                string `json:"Make"`
//		Model               string `json:"Model"`
//		Id string `json:"Idd"`
//	}
//type Test struct {
//	Count string `json:"count"`
//}

func main() {
	//url := "http://127.0.0.1:8888/api/app/ticket/create" //前面没有http报错 first path segment in URL cannot contain colon
	//url := "http://accidentsrecordstest.wiltechs.com:8089/api/accident/vehicles"
	//data := map[string]interface{}{
	//	"caseId": "1120202210521478992",
	//}
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println("JSON 编码错误：", err)
	//	return
	//}
	//req, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	fmt.Println("创建http请求错误 : ", err)
	//	return
	//}
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer tFVLRskE2zDZXhYDscJmaI3xTk2qhUtWl8JjvLomSiEEqwvZm02kgTRLlUwG2JBm-zeWaV_yDN4eAajCOZ6ptmJuLYHZ015yYF6WU8vz1JpklzUwT63ye5w9BYpi1zFa6ZjyNOZKEv95sZ6eqsYvAJDF-pxfBuKRUq_0ZdoueBKaREscwYvSci2PrRQt6PqPZkLJT7lX7ctURp5SwraTesYdQdYR9Q8tC1iIwuF63zpoDCiQJjZ8P5jkNm1rI-wXFury8tf274DBPGjNls6YKOcsL1_eOgCSuPkIIQUg-8yKc3zMrszBBfgpbHOCdeEmXtK1uO7H0AJhlfXNXrz7me4p9Xo")
	//
	//client := &http.Client{Timeout: 10 * time.Second}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("HTTP 请求错误：", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("响应状态码：", resp.StatusCode)
	//fmt.Println("响应内容：")
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("读取响应内容错误：", err)
	//	return
	//}
	//fmt.Println(string(body))
	//
	//var cars []Car
	//err = json.NewDecoder(resp.Body).Decode(&cars)
	//if err != nil {
	//	fmt.Println("JSON 解析失败:", err)
	//	return
	//}
	//for _, car := range cars {
	//	fmt.Printf("ID: %d, Model: %s, Number: %s\n", car.Id, car.Model, car.VehicleSerialNumber)
	//}
	url := "http://accidentsrecordstest.wiltechs.com:8089/api/accident/vehicles"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建http请求错误 : ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer tFVLRskE2zDZXhYDscJmaI3xTk2qhUtWl8JjvLomSiEEqwvZm02kgTRLlUwG2JBm-zeWaV_yDN4eAajCOZ6ptmJuLYHZ015yYF6WU8vz1JpklzUwT63ye5w9BYpi1zFa6ZjyNOZKEv95sZ6eqsYvAJDF-pxfBuKRUq_0ZdoueBKaREscwYvSci2PrRQt6PqPZkLJT7lX7ctURp5SwraTesYdQdYR9Q8tC1iIwuF63zpoDCiQJjZ8P5jkNm1rI-wXFury8tf274DBPGjNls6YKOcsL1_eOgCSuPkIIQUg-8yKc3zMrszBBfgpbHOCdeEmXtK1uO7H0AJhlfXNXrz7me4p9Xo")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP 请求错误：", err)
		return
	}
	defer resp.Body.Close()
	//var cars []Car
	//var test []Test
	//var t []string
	var todos []map[string]interface{}
	//t := []string{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	//for _, car := range cars {
	//	fmt.Printf("ID: %s, Model: %s, Number: %s\n", car.Id, car.Model, car.VehicleSerialNumber)
	//}
	//for _, car := range cars {
	//	fmt.Printf("ID: %s\n", car.Id)
	//}
	fmt.Println(todos)
	fmt.Println(len(todos))

}
