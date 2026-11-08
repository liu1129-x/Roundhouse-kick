package main

import "fmt"

func Todel() {
	var name, hemo string
	var age, tizhogn, shengao int
	fmt.Print("请输入姓名、年龄、身高、体重:")
	fmt.Scanf("%s %d %d %d", &name, &age, &tizhogn, &shengao)
	// name = "王兴"
	// hemo = "China"
	// age = 21
	// tizhogn = 65
	// shengao = 173
	fmt.Print("姓名: ", name, "\n", "地址: ", hemo, "\n", "年龄: ", age, "\n",
		"体重: ", tizhogn, "\n", "身高: ", shengao, "\n")
}
