package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	a := decimal.NewFromFloat(1.12)
	b := decimal.NewFromFloat(1.4)
	c := decimal.NewFromFloat(2.2)

	sum := decimal.Sum(a, b, c)

	d := 1.12
	e := 1.4
	f := 2.2
	fmt.Println(d + e + f)

	fmt.Println(sum) // 输出: 4.72
}
