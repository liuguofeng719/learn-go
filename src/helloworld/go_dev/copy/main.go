package main

import (
	"fmt"
)

func main() {
	var num1 = []int{3, 5, 6, 7, 8}
	var num2 = make([]int, 1)
	copy(num2, num1)
	fmt.Println("size=", len(num2))
	fmt.Println("cap=", cap(num2))
	for i := 0; i < len(num2); i++ {
		fmt.Printf("%v", num2[i])
	}
}
