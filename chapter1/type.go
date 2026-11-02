package main

import (
	"fmt"
)

type Person struct {
	FirestName string
	LastName   string
	Age        int
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

type Operation func(int, int) int

func main() {
	fmt.Println("整数类型(Integer Type)")
	var x int = 10
	var y int64 = 100

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("浮点数类型(Floating-Point Types)")
	var a float32 = 3.14
	var b float64 = 3.14159265359

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("布尔类型(Boolean Type)")
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue)
	fmt.Println(isFalse)

	fmt.Println("字符串类型(String Type)")
	str1 := "Hello,"
	str2 := "Go!"

	concatenated := str1 + str2
	fmt.Println(concatenated)

}
