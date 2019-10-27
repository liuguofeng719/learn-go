package main

import (
	"log"
)

/*
	1、函数声明
		func 函数名称(参数名 参数类型) 返回值类型 {}  一个返回值
		func 函数名称(参数名 参数类型) (返回值类型，返回值类型) {} 多个返回值
	2、函数调用
	3、匿名函数
	4、函数多返回值
	5、闭包
	6、变量函数
	7、参数传递函数
	8、通过函数名的大小写来暴露访问级别，小写只能在本包中可以访问，大写可以跨包访问
	9、函数没有重载，重写
	10、函数可变参数
*/
var x, y = 1, 2

func main() {

	allFunc()

	var f1 = func() {
		log.Println("我是匿名函数")
	}

	log.Printf("%p,%v,%T\n", f1, &f1, f1)

	f1()
	func() {
		log.Println("xxxx")
	}()

	func(x int) {
		log.Println(x)
	}(10)

	func(x int, y int) {
		log.Println(x * y)
	}(10, 20)

	func() {
		log.Println(x * y)
	}()
}

func allFunc() {
	show()
	show1(10, 20)
	show2(10, 20)
	var sum = show3(10, 20)
	log.Println(sum)
	a, b := show4(10, 20)
	log.Println(a, b)

	show5(func() {
		log.Println("5555")
	})
	show6(func(a int, b string) {
		log.Printf("%d,%s\n", a, b)
	})

	var c = show7(func(a, b string) {
		log.Println(a, b)
	})
	c()

	var f1 = show8()
	a1, b1 := f1(0, "go6")
	log.Println(a1, b1)
	a2, b2 := f1(1, "go6")
	log.Println(a2, b2)

	show9(1, 2, 3, 4, 5)
}

func show() {
	log.Println("show() === ")
}

func show1(a int, b int) {
	log.Println(a, b)
}

func show2(a, b int) {
	log.Println(a, b)
}

func show3(a int, b int) int {
	return a + b
}

func show4(x, y int) (int, int) {
	return x * 10, y * 10
}

// 传递函数
func show5(f func()) {
	f()
}

// 传递函数带参数
func show6(f func(int, string)) {
	f(10, "块钱")
}

// 带函数返回值的函数
func show7(f func(string, string)) func() {
	f("xx", "yyy")
	return func() {
		log.Printf("%s\n", "哈哈装逼失败")
	}
}

// 返回函数且函数参数为int，string且返回值也是(int,string)
func show8() func(int, string) (int, string) {
	log.Println("show8")
	var count = 0
	return func(a int, b string) (int, string) {
		count++
		return count, b
	}
}
func show9(b ...int) {
	for _, b2 := range b {
		log.Println(b2)
	}
}
