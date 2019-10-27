package main

import "fmt"

func main() {
	// lable2:
	for i := 0; i < 4; i++ {
	lable1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break lable2
				break lable1
			}
			fmt.Printf("i=%d,j=%d \n", i, j)

		}
	}
}
