package main

import "fmt"

func main() {
	bmi()
}

// bmi函数
func bmi() {
	//年龄
	var age int
	fmt.Print("请输入年龄:")
	fmt.Scan(&age)
	//身高
	var hear float64
	fmt.Printf("请输入身高(m):")
	fmt.Scan(&hear)
	//体重
	var d float64
	fmt.Printf("请输入体重(kg):")
	fmt.Scan(&d)

	//计算公式
	bmi := d / (hear * hear)
	fmt.Printf("BMI值为:%.1f\n", bmi)

	//bmi值判断
	if 23.9 >= bmi && bmi >= 18.5 {
		fmt.Println("正常")
	} else if 27.9 >= bmi && bmi >= 24 {
		fmt.Println("超重")
	} else if bmi >= 28 {
		fmt.Println("肥胖")
	} else if 18.5 > bmi && bmi > 0 {
		fmt.Println("偏瘦")
	} else {
		fmt.Println("输入错误！")
	}
}
