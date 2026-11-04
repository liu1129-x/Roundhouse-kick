package main

import "fmt"

func ZhengShu(x int, y int) int {
	return (x + y) / 2
}

func main() {
	var a, b int
	fmt.Printf("请输入两位整数,用空格分开:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("值:%d\n", ZhengShu(a, b))
}
