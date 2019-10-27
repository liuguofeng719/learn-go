package main

import "fmt"

func main() {
	// 切片对字符串操作，
	var str = "abc@163.com"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	newstr := str[3:] // 163.com
	fmt.Printf("%v", newstr)

	// 如果有中文不转换会出现乱码，解决方案2种
	// 1、for i,v:=range var {} 循环
	// 2、把字符串转换成[]rune(str1)
	var str1 = "abc@163.com网易"
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i])
	}
	newstr1 := str2[3:] // 163.com
	fmt.Printf("%c\n", newstr1)

	fbn(10)
}

// 斐波拉契数列
// 1 1 2 3 5 8 13 21
func fbn(n int) []int64 {
	num := make([]int64, n)
	num[0] = 1
	num[1] = 1
	for i := 2; i < n; i++ {
		num[i] = num[i-1] + num[i-2]
	}
	return num
}
