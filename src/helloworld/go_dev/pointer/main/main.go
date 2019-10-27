package main

import "log"

func main() {
	var ii int = 1
	log.Printf("%p", &ii)

	var intPtr *int
	intPtr = &ii
	log.Printf("%p,%v,%v", intPtr, &intPtr, *intPtr)
}
