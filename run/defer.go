package main

import "fmt"

func main() {
	fmt.Println(returnButDefer())
}

func returnButDefer() (t int) {
	defer func() {
		t = t * 10
	}()

	return 1
}
