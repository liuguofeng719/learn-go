package main

import "fmt"

func main() {
	// 除法运算符 /
	// 如果运算的数都是整数，那么除后，会去掉小数部分，保留整数部分
	fmt.Println(10 / 4) // 2
	var div float32 = 10 / 4
	fmt.Println(div) // 2
	// 如果要保留小数部分，运算的数必须带小数部分
	var div1 float32 = 10.0 / 4
	fmt.Println(div1) // 2.5

	// 取模运算符 %
	// 计算公式 a % b = a - a / b * b
	fmt.Println("-10%3=", -10%3)   // -1
	fmt.Println("10%-3=", 10%-3)   // 1
	fmt.Println("-10%-3=", -10%-3) // -1

	var a int = 10
	var b int = 20
	a = a + b // a = 30
	b = a - b // b = 10
	a = a - b // a = 20
	fmt.Println("a=", a, "b", b)
	var aa, bb = 20, 10
	fmt.Println("aa=", aa, "bb=", bb)

}
