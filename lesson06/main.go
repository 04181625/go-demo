package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义
	//解析
	//渲染模版
	t, err := template.New("f").ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "汉堡包"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/hello", f1)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
		return
	}
}
