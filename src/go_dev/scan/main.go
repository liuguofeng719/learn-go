package main

import (
	"fmt"
)

func main() {
	// 通过控制台接收用户信息【姓名、年龄、薪水、是否考试通过】
	// var name string
	// var age byte
	// var sal float32
	// var isPass bool
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)
	// fmt.Println("是否考试通过：")
	// fmt.Scanln(&isPass)

	// fmt.Printf("姓名 %v \n 年龄 %v \n 薪水 %v \n 是否通过 %v \n", name, age, sal, isPass)

	// fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	// fmt.Printf(" 姓名 %v \n 年龄 %v \n 薪水 %v \n 是否通过 %v \n", name, age, sal, isPass)
	// var i int
	// var j int
	// j := i++
	// fmt.Println("%d",i)

	var b bool = true
	if b == false {
		fmt.Println("false")
	} else if b {
		fmt.Println("b")
	}
}
