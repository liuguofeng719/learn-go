package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)

	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-ticker
	}

	abort := make(chan struct{}, 1)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-abort:
		fmt.Println("xxxx")
	case <-time.After(time.Second):
		fmt.Println("after'")
	}
}
