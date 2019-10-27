package main

import (
	"fmt"
)

func main() {
	arrays()
}

func arrays() {
	var tmp [3]string
	tmp[0] = "张三"
	tmp[1] = "李四"
	tmp[2] = "王五"

	var strs = [3]string{"a", "b", "c"}

	fmt.Println(strs)
	fmt.Println(tmp)

	for k, v := range tmp {
		fmt.Println(k, v)
	}
}
