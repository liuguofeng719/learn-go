package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	testFunc()
	go running()
	// 让程序等待
	var inputString string
	fmt.Scanln(&inputString)
}

func TestRouting(a int) {
	fmt.Println(a)
}

func testFunc() {
	for index := 0; index < 100; index++ {
		go TestRouting(index)
	}
	time.Sleep(time.Second)
}

func running() {
	var times = 0
	for {
		times++
		log.Printf("滴滴%v", times)
		time.Sleep(time.Second)
	}
}
