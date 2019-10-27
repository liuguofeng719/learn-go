package main

import "fmt"

func main() {
	var arr = [...]int{1, 3, 8, 10, 44, 56, 100}
	binarySort(&arr, 0, len(arr), 100)
	// 把递归变成循环
	whileSort(&arr, len(arr), 8)
}

func whileSort(arr *[7]int, hight int, findVal int) (b *[7]int) {

	return arr
}

// 二分查找，就是折半查找
// middeIndex :=  (leftIndex + rightIndex) / 2
// 1、arr[middle] > findVal 说明查找的元素在左边，binarySort(arr,leftIndex,middle-1,findVal)
// 2、arr[middle] < findVal 说明查找的元素在右边，binarySort(arr,middle+1,rightIndex,findVal)
// 3、leftIndex > rightIndex 说明元素没有找到
func binarySort(arr *[7]int, leftIndex int, rightIndex int, findVal int) {
	// 这里肯出现越界，所以需要优化，解决方案leftindex+ (RightIndex - leftIndex) / 2 注意运算符的优先级
	middleIndex := leftIndex + (rightIndex-leftIndex)/2
	// 这里除以2可以使用位移右位移，也需要注意运算符的优先级
	// 方案 middelIndex := leftIndex + ((rightIndex- leftIndex) >> 1)
	if leftIndex > rightIndex {
		fmt.Println("没有找到该元素", findVal)
		return
	}
	if arr[middleIndex] > findVal {
		binarySort(arr, leftIndex, middleIndex-1, findVal)
	} else if arr[middleIndex] < findVal {
		binarySort(arr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Println("查找的元素索引为", middleIndex)
	}
}
