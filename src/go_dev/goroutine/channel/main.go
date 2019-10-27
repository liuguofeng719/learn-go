package main

import (
	"fmt"
	"log"
	"time"
)

// 通道使用,通道声明必须配合make才能使用
// var 通道变量 chan 通道类型
// 通道实例 := make(chan 数据类型)
// 通道没有消费者，会报错 all goroutines are asleep - deadlock!
// 接收数据有2中方式 阻塞方式 data:=<-ch，
// 异步方式 data,ok:=<-ch
// 非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行
func main() {
	// ch3 := make(chan *Equal) //声明并初始化结构体通道
	// ch3 <- &Equal{eq: "xxx"}
	// time.Sleep(time.Second)
	// c3, ok := <-ch3
	// if ok {
	// 	log.Println(c3.eq)
	// }
	var c ByteCounter
	c.Write([]byte("hello"))
	log.Printf("%v",c)

	ch4:=make(chan int ,10)
	ch4 <- 2
}

type ByteCounter int 
func (b *ByteCounter) Write(p []byte) (n int ,err error){
	*b += ByteCounter(len(p))
	return len(p),nil
}
type Equal struct {
	eq string
}

func test_chan_int() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 12
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	name := <-ch
	fmt.Println("all done", name)
}

func test_chan_interface() {
	ch2 := make(chan interface{}) // 声明并初始化任意类型通道
	go func() {
		for {
			ch2 <- "sss"
		}
	}()

	for {
		name := <-ch2
		log.Println(name)
		time.Sleep(time.Second)
	}
}
