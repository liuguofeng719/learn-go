package main

import "fmt"

func main() {
	// var b1 map[int]int
	// var b2 map[string]string
	// var b3 map[bool]bool
	// var b4 map[string]bool
	// var b5 map[string]int
	// var b6 map[string]interface{}
	// var b7 map[float32]float32
	// var b8 map[Student]string
	// var b9 []map[string]string
	var b10 map[string]map[string]string
	var b11 map[string][]string
	var b12 map[interface{}][]string
	// var b13 map[interface{}]func
	fmt.Println(b10, b11, b12)
	var b13 = make(map[interface{}][]string)
	b13[1] = []string{"33", "33", "44"}
	b13["1"] = []string{"33", "33", "44"}
	b13[true] = []string{"33", "33", "44"}
	for k, v := range b13 {
		fmt.Printf("k = %v,v = %v \n", k, v)
	}
	// 声明切分配空间为初始空间为3
	var b14 = make(map[string]string, 3) //
	b14["a"] = "123"
	b14["b"] = "456"
	b14["c"] = "567"
	b14["d"] = "789"
	fmt.Println(len(b14))
	for k, v := range b14 {
		fmt.Printf("k = %v,v = %v \n", k, v)
	}
	// 第三种声明使用方式
	var b15 = map[string]string{
		"a1": "a1",
		"a2": "a2",
	}
	for k, v := range b15 {
		fmt.Printf("k = %v,v = %v \n", k, v)
	}

	delete(b15, "a1")
	fmt.Println("=====")
	for k, v := range b15 {
		fmt.Printf("k = %v,v = %v \n", k, v)
	}
	b15["a2"] = "xx"
	fmt.Printf("k = %v\n", b15["a2"])

}

type Student struct{}
