package main

import "fmt"

type Usb interface {
	Start()
	End()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("iphone start")
}
func (p Phone) End() {
	fmt.Println("iphone End")
}
func (p Phone) Call() {
	fmt.Println("打电话")
}

type Mouse struct {
}

func (p Mouse) Start() {
	fmt.Println("Mouse start")
}
func (p Mouse) End() {
	fmt.Println("Mouse End")
}

type Computer struct {
}

// 通过类型断言来执行不同的对象里面的方法
// 这里和java中 instanceof 一样的功能
func (c Computer) Working(usb Usb) {
	usb.Start()

	if phone, ok := usb.(Phone); ok {
		fmt.Printf("ok=%v\n", ok)
		phone.Call()
	}
	usb.End()
}

func main() {
	var c = Computer{}
	var p = Phone{}
	var m = Mouse{}
	c.Working(p)
	c.Working(m)

	var ps interface{}
	ps = 12
	fmt.Printf("%v,%T\n", ps, ps)
	// b = ps 会报错
	// cannot use ps (type interface {}) as type int in assignment: need type assertion
	var b int
	// b = ps
	b = ps.(int) // 使用需要、类型断言
	fmt.Printf("%v,%T\n", b, b)

	parameter(1)
	parameter("2")
	parameter(1.1)
}

func parameter(data interface{}) {
	switch _type := data.(type) {
	case int:
		fmt.Printf("%v,%T\n", _type, _type)
	case float32, float64:
		fmt.Printf("%v,%T\n", _type, _type)
	case string:
		fmt.Printf("%v,%T\n", _type, _type)
	default:
		fmt.Println("未知类型")
	}
}
