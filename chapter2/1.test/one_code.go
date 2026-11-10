package main

import "fmt"

func main() {
	calculator()
}

func calculator() {
	var x, y int  //定义x y两个整数变量
	var op string //对应 + - * /
	fmt.Print("请输入表达式(例如:10 + 5):")
	//判断返回值
	n, err := fmt.Scanf("%d%s%d", &x, &op, &y)
	// 同于判断输入格式错误
	if err != nil || n != 3 {
		fmt.Println("输入格式错误!")
		return
	}

	switch op {
	case "+":
		fmt.Println("和:", x+y)
	case "-":
		fmt.Println("差:", x-y)
	case "*":
		fmt.Println("乘:", x*y)
	case "/":
		if y != 0 {
			fmt.Println("除:", x/y)
		} else {
			fmt.Println("除数不能为零！")
		}
	default:
		fmt.Println("输入错误！")
	}
}
