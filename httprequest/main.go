package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://127.0.0.1:8888/api/app/ticket/create" //前面没有http报错 first path segment in URL cannot contain colon
	data := map[string]interface{}{
		"caseId":            "99999",
		"category":          "diahidnait",
		"caseStatus":        2,
		"badge":             "yp1234567899",
		"workStatus":        3,
		"injuryDate":        "2022-11-20",
		"injuryBodyPart":    "head",
		"injuryDescription": "aa爆炸",
		"wareHouseCode":     "105214",
		"staffId":           "35114",
		"createBy":          "1234",
		"updateBy":          "213",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON 编码错误：", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建http请求错误 : ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzIiwidXNlcl9uYW1lIjoic2FsbHkubCIsInN0YWZmX2lkIjoiM2NlMzhhNmEtZWIyMy00YjJmLWJlYzMtNTZlNzYwNzk3Njk2Iiwicm9sZV90eXBlIjowLCJleHAiOjE2Nzc2NjAzNzUsImlzcyI6ImluanVyZWQgcGxhdGZvcm0gc2VydmljZSJ9.3vUQKFmY8bRtX7VQdhlsobXqNZvqIfjF3dHJCbDUCKY")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP 请求错误：", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("响应状态码：", resp.StatusCode)
	fmt.Println("响应内容：")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容错误：", err)
		return
	}
	fmt.Println(string(body))
}
