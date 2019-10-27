package main

import "fmt"

func test(key byte) byte {
	return key
}

// swith case 表达式支持任意类型匹配
func main() {
	var num int = 20
	switchNum(num)
	switchByte('a')
	switchString("b")
	switchFunc('b')

	var num1 int16 = 3333
	switchType(num1)
	switchFallthough('a')
}

// switch fallthough
func switchFallthough(key byte) {
	switch test(key) {
	case 'a':
		fmt.Println()
		fmt.Println("=====a")
		fallthrough // 默认穿透一层s
	case 'b':
		fmt.Println("=====b")
	case 'c':
		fmt.Println("===c")
	case 'd':
		fmt.Println("===d")
	default:
		fmt.Println("unkown")
	}
}

//
func switchType(key interface{}) {
	switch _type := key.(type) {
	case byte:
		fmt.Printf("byte %T, %v", _type, _type)
	case string:
		fmt.Printf("%T, %v", _type, _type)
	case int8, int16, int, int32, int64:
		fmt.Printf("int - %T, %v", _type, _type)
	case nil:
		fmt.Printf("%T, %v", _type, _type)
	case float32, float64:
		fmt.Printf("%T, %v", _type, _type)
	case func(int):
		fmt.Println("func")
	case bool:
		fmt.Printf("%T, %v", _type, _type)
	default:
		fmt.Println("unkown")
	}
}

// switch 函数
func switchFunc(key byte) {
	switch test(key) {
	case 'a':
		fmt.Println("a")
	case 'b':
		fmt.Println("b")
	case 'c':
		fmt.Println("c")
	case 'd':
		fmt.Println("d")
	default:
		fmt.Println("unkown")
	}
}

// 字符串
func switchString(key string) {
	switch key {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c'":
		fmt.Println("c")
	case "d":
		fmt.Println("d")
	default:
		fmt.Println("unkown")
	}
}

// byte case
func switchByte(key byte) {
	switch key {
	case 'a':
		fmt.Println("a")
	case 'b':
		fmt.Println("b")
	case 'c':
		fmt.Println("c")
	case 'd':
		fmt.Println("d")
	default:
		fmt.Println("unkown")
	}
}

// 数字比较
func switchNum(num int) {
	switch num {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("default")
	}
}
