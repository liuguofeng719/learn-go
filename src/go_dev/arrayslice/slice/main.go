package main

import (
	"fmt"
)

func main() {

	var num = new(int)
	fmt.Printf("num=%T\n", num)

	// 声明切的3中方式
	var arrays = [5]int{3, 10, 30, 2, 40}
	// 数组使用切片
	slice1 := arrays[1:] // 10, 30, 2, 40
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("值=%v,类型=%T,指针地址=%p\n", slice1[i], slice1[i], &slice1[i])
	}

	// 切片在切片
	slice2 := slice1[1:3] // 30, 2
	for i, v := range slice2 {
		fmt.Printf("index= %v,value = %v \n", i, v)
	}

	// 声明切片
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	// 使用make声明切片
	var slice3 = make([]int, 4, 4)
	slice3[0] = 20
	slice3[1] = 30
	fmt.Println(slice3)
	// 使用append动态增加切片的值
	slice3 = append(slice3, 10, 11, 30)
	slice = append(slice, slice3...)

	for index, val := range slice {
		fmt.Printf("index=%v,value=%v\n", index, val)
	}
}
