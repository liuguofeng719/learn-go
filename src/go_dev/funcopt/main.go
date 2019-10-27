package main

import (
	"fmt"
	"strconv"
	"strings"
)

var myf func(name string, age int) string

func main() {
	strs := show(sayHello)
	fmt.Println(strs)
	myf = addUser("mm", 20)
	xx := myf("xx", 21)
	fmt.Println(xx)

	// 函数变量
	var show = func(age int, address string) string {
		age1 := fmt.Sprintf("%d", age)
		return address + age1 + strconv.Itoa(age)
	}
	fmt.Println(show(22, "chengdushi"))

	// 匿名函数
	test1 := func(age int, idNo string) string {
		return idNo + strconv.Itoa(age)
	}(20, "232443432")

	fmt.Println(test1)

	// 闭包
	f := accumulate(1)
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	if true {
		var b = 10
		if true {
			var b = 10
			fmt.Println(b)
		}
		fmt.Println(b)
	}
	num, _ := strconv.Atoi("123")
	fmt.Println("num=", num)
	index := strings.Index("hello", "l")
	fmt.Println(index)
}
func show(myfunc func(string, int) string) string {
	return myfunc("zhangsan", 20)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("我叫：%s,年龄：%d", name, age)
}

func addUser(name string, age int) func(string, int) string {
	return func(name1 string, age1 int) string {
		return fmt.Sprintf("%s,%d", name+name1, age+age1)
	}
}
func accumulate(ac int) func(int) int {
	return func(acc int) int {
		ac = ac + acc
		return ac
	}
}
