package greetings

// package main

import (
	"fmt"
	// "golang.org/x/text/message"
)

// func Hello(name string) string {
// 	message := fmt.Sprintf("hello,%v,wlcome!", name)
// 	return message
// }

// func Sum(x int) int {
// 	// var message int
// 	fmt.Print("155", x)
// 	return 8421
// }

// func tom1(m int) int {
// 	// fmt.Println(m)
// 	return 10 + m
// }

// func Tom(a int, p *int) int {
// 	// m = 100
// 	// result := tom1(m)
// 	// m = 100

// 	a = 20
// 	*p = 30
// 	return a + *p
// }

// func main() {
// 	fmt.Println(Tom(1))
// }

type User struct {
	Name string
	Age  int
	ID   int
}

func Peloep() {
	User1 := User{Name: "Tom", Age: 22, ID: 1001}
	fmt.Println(User1)
}
