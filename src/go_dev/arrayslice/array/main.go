package main

import "fmt"

func main() {
	// 声明intArr数组，有3个元素，数组的所有从0开始
	var intArr [3]int // 默认值[0 0 0]
	// 默认值为int数组、float 为0，string为空串，bool 为false
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	// 地址分配规则 = 首个元素的地址+数组类型的长度 &intArr + 8（因为int是八个字节）
	// &intArr 就是首个元素的地址
	fmt.Printf("first=%p,one=%p,two=%p,three=%p \n", &intArr, &intArr[0], &intArr[1], &intArr[2])
	// 内存地址
	// first=0xc0000b6000,one=0xc0000b6000,two=0xc0000b6008,three=0xc0000b6010

	test1(intArr)
	fmt.Println("数组值是否发生改变", intArr)

	test2(&intArr)
	fmt.Println("数组值是否发生改变", intArr)
	iteatorByte()
	var nums = [5]int{3, 2, 5, 4, 1}
	twoNumMax(nums)
	var num1 = [...]int{1, 0, 5, 2, 9}
	twoNumMax(num1)
	// 类型推导
	var num2 = [...]int{1: 9, 2: 6, 0: 20, 3: 10, 4: 2}
	num3 := [...]int{1: 9, 2: 6, 0: 20, 3: 10, 4: 20}
	twoNumMax(num2)
	twoNumMax(num3)
}

func test1(ar [3]int) {
	ar[2] = 20
}

func test2(ar *[3]int) {
	ar[2] = 20
}

func iteatorByte() {
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i) // 'A' + 1 = 'B' 以此类推
	}

	for i, v := range myChar {
		fmt.Printf("index = %d,value = %c \n", i, v)
	}
}

// 找出数组中最大的那个数
func twoNumMax(nums [5]int) {
	var maxVal = 0
	var maxIndex = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex++
		}
	}
	fmt.Printf("maxVal = %v,maxIndex = %v", maxVal, maxIndex)
}
