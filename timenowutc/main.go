package main

import (
	"fmt"
	"time"
)

func main() {
	now1 := time.Now().UTC()
	fmt.Println(now1)
	fmt.Println("===========")
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("error")
	}
	now2 := time.Now().In(location)
	fmt.Println(now2)
	fmt.Println("===========")
	location1, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("error")
	}
	now3 := time.Now().In(location1)
	fmt.Println(now3)
	fmt.Println("===========")
	current := time.Now().UTC()
	fmt.Println(current.Format("2006-01-02 15:04:05"))
}
