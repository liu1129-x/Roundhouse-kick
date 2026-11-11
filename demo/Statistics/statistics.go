package main

import (
	"fmt"
)

func main() {
	Menu()
}

const handwriting = 5.0 //卷面分
const serious = 6.0     //仔细分（认真分）

// 输入菜单
func Menu() {
	var math, chinese, english float32

	// var name string
	// fmt.Printf("姓名:")
	// fmt.Scanf(&name)

	fmt.Printf("请输入数学、语文、英语分数(用空格分开):")
	fmt.Scanf("%f %f %f", &math, &chinese, &english)

	fmt.Printf("总分为:%.1f\n", Sum(math, chinese, english))
	fmt.Printf("平均分为:%.1f\n", Average(math, chinese, english))

}

// 净分值
func Totalscore(math ...float32) float32 {
	var sum float32 = 0.0

	for i := 0; i < len(math); i++ {
		sum += math[i] //分数进行累加
	}

	if sum >= 270.0 {
		fmt.Println("甲等")
	} else if sum >= 240 {
		fmt.Println("乙等")
	} else if 240 > sum && sum >= 200 {
		fmt.Println("丙等")
	} else if 200 > sum && sum >= 180 {
		fmt.Println("丁等")
	} else {
		fmt.Println("留级")
	}
	return sum
}

// 总分
func Sum(math ...float32) float32 {
	var sum float32 = 0.0

	for i := 0; i < len(math); i++ {
		sum += math[i] //分数进行累加
	}

	return sum //返回和
}

// 平均分
func Average(math ...float32) float32 {
	var sum float32 = 0.0

	for i := 0; i < len(math); i++ {
		sum += math[i]
	}

	return (float32(sum)) / (float32(len(math)))
}
