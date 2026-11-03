package main

import "fmt"

func ZhengShu(x int, y int) int {
	return (x + y) / 2
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", ZhengShu(a, b))
}
