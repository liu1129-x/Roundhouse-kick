package main

import "fmt"

func main() {
	var math, chinese, english float64
	fmt.Printf("请输入三科科目分数,用空格分开:")
	fmt.Scanf("%f %f %f", &math, &chinese, &english)
	fmt.Printf("总分为:%.1f\n", Sum(math, chinese, english))
	fmt.Printf("平均分为:%.1f\n", Average(math, chinese, english))
}

// 总分
func Sum(math ...float64) float64 {
	sum := 0.0

	for i := 0; i < len(math); i++ {
		sum += math[i] //分数进行累加
	}

	return sum //返回和
}

// 平均分
func Average(math ...float64) float64 {
	sum := 0.0

	for i := 0; i < len(math); i++ {
		sum += math[i]
	}

	return (float64(sum)) / (float64(len(math)))
}
