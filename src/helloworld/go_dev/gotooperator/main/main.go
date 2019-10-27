package main

import (
	"fmt"
	"helloworld/go_dev/gotooperator/utils"
)

func main() {
	var num int = 20
	fmt.Println("ok1")
	if num > 10 {
		goto label1
	}
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
label1:
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	utils.ConvertStr()
}
