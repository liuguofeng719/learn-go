package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	one := 1
	one, two := 2, 3
	fmt.Println(one, two)
	if true {
		one := 4
		fmt.Println(one)
	}

	// 声明一个缓冲大小为1，且阻塞队列
	var intChan = make(chan int, 1)
	intChan <- 3

	close(intChan)

	// 如果不close 取出intChan里面的值，就等待管道里面写入值
	// 这里生产者不会生成数据，这里就会出现死锁
	for k := range intChan {
		fmt.Println(k)
	}

	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		nR := make(chan int64)
		nW := make(chan int64)
		go natualR(nR)
		go squareR(nW, nR)

		for k := range nW {
			fmt.Println(k)
		}
		group.Done()
	}()
	group.Wait()
}

func caseOnlyRW() {
	// 只写通道 chan<- int
	writeOnlyChannel := make(chan<- int, 5)
	writeOnlyChannel <- 1

	// 只读通道 <-chan int
	readOnlyChannel := make(<-chan int, 5)
	//readOnlyChannel <- 2
	<-readOnlyChannel // 读取到值就扔掉
}

func natualR(natualR chan<- int64) {
	for i := 0; i < 5; i++ {
		natualR <- int64(i)
	}
	close(natualR)
}

func squareR(out chan<- int64, in <-chan int64) {
	for {
		o, ok := <-in
		if !ok {
			break
		}
		out <- o * o
	}
	close(out)
}

// 通道当着用管道
func pip() {
	natural := make(chan int64)
	result := make(chan int64)

	go naturals(natural)
	go squares(natural, result)
	for v := range result {
		log.Println(v)
	}
}

func naturals(natural chan int64) {
	for i := 0; i < 5; i++ {
		natural <- int64(i)
	}
	// 关闭通道
	close(natural)
}

func squares(natural chan int64, result chan int64) {
	for {
		v, ok := <-natural
		// 读取去完，跳出循环
		if !ok {
			log.Println("自然数已经读取完成")
			break
		}
		result <- v * v
	}
	// 关闭通道
	close(result)
}
