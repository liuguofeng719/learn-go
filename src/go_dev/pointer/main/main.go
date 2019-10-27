package main

import (
	"fmt"
	"log"
)

func main() {
	var ii int = 1
	log.Printf("%p", &ii)

	var intPtr *int
	intPtr = &ii
	log.Printf("%p,%v,%v", intPtr, &intPtr, *intPtr)

	var num int8 = 10
	fmt.Printf("%p,%v", &num, num)

	var num1 *int8
	num1 = &num
	*num1 = 20
	fmt.Printf("%p,%v", num1, *num1)
}
