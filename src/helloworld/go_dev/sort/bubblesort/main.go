package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{20, 5, 10, 30, 60}
	bubbleSort(&arr)
}

func bubbleSort(arr *[5]int) {
	// 为什么减一，剩下最后一个，不需要比较了
	for i := 0; i < len(*arr)-1; i++ {
		var temp int
		// 每执行一轮都会比较出最大的元素，外层循环在来，不需要全部元素都比较
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	for _, v := range *arr {
		fmt.Printf("value = %v ", v)
	}
}
