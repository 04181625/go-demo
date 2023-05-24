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

	a := "2023-05-09 05:56:39"
	b, err := time.Parse("2006-01-02 15:04:05", a)
	if err != nil {
		fmt.Println(err)
	}
	pstTimezone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Failed to load PST timezone:", err)
		return
	}
	fmt.Println(b.In(pstTimezone))
}
