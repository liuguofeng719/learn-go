package main

import "fmt"

// 通过全局方式声明管道
var pipe chan int

func main() {
	pipe = make(chan int, 1)
	go go_add(10, 10)

	// 通过参数传入管道
	var p chan int
	p = make(chan int ,1)

	go go_add2(20,20,p)

	sum := <-pipe

	fmt.Println("====", sum)
}

func go_add(a int, b int) {
	var sum int
	sum = a + b

	pipe <- sum
}

func go_add2(a int ,b int ,c chan int) {
	var sum int
	sum  = a+ b
	c <- sum
}

func go_testroute(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func gotest1() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
