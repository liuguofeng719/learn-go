package main

import (
	"fmt"
	_ "helloworld/go_dev/day2/test"
	"os"
)

var cwd string
var err error

func init() {
	cwd, err = os.Getwd()
	if err != nil {
		fmt.Printf("%s,%s", cwd, err)
	}
}

func main() {
	var cwd, err = os.Getwd()
	if err != nil {
		fmt.Printf("%s,%s", err, cwd)
	}
	fmt.Printf("%s", err)
	// var pipe chan int = make(chan int, 1)
	var pipe = make(chan int, 3)
	pipe <- 2
	var t1 int
	t1 = <-pipe
	fmt.Println(t1)
}
