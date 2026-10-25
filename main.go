package main

import "fmt"

func main() {
	bmi()
}

func bmi() {
	//年龄
	age := 18
	//身高
	h := 1.72
	//体重
	d := 65.0

	fmt.Print("年龄:", age)
	fmt.Print("\n")
	fmt.Printf("BMI值:%.1f\n", d/(h*h))

}
