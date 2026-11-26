package main

import (
	"fmt"
)

var a1 int = 7
var b1 int = 9

func main() {
	a1, b1, c1 := 10, 20, 0
	fmt.Println("a1 = ", a1)
	fmt.Println("b1 = ", b1)
	fmt.Println("c1 = ", c1)
	fmt.Println("----------")
	sum(a1, b1)
}

func sum(a1, b1 int) (c1 int) {
	a1++
	b1 += 2
	c1 = a1 + b1
	fmt.Println("a1 = ", a1)
	fmt.Println("b1 = ", b1)
	fmt.Println("c1 = ", c1)
	return c1
}
