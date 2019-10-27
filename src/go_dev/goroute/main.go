package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	go running()
	// 让程序等待
	var inputString string
	fmt.Scanln(&inputString)
}

func test() {
	for index := 0; index < 100; index++ {
		go go_route(index)
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
