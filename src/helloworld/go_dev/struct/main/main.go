package main

import "log"

/*
	结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存
	结构体在实例化后会形成指针类型的结构体

	实例化结构体方式有3中
	1、var 实例名称 T类型  T 代表结构体的类型 var p = Point
	2、通过var 实例名称 = new(结构体类型) var p = new(Point)
	3、通过&取地址符  var ins = &Point{}

	4、匿名结构体
		ins := struct{
			字段名称 字段类型
		}{
			字段名称:字段的值
		}
		ins := &struct{
			字段名称 字段类型
		}{
			字段名称:字段的值
		}
	5、为结构体添加方法
	type Person struct {
		name string
		age int
	}
	接收器分为指针类型接收器和普通类型的接收器
	指针类型用于大对象传递，复制指针速度快，小对象用普通类型传递
	func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    	函数体
	}
	func (p Person) getPerson() Person {

	}
	ins := new(Person)
	p:=ins.getPerson()

	6、为任意类型添加方法
	type MyInt int
	func (myInt MyInt) IsZero() bool {
		return myInt == 0
	}
	var m MyInt
	fmt.Println(m.IsZero())

	7、结构体嵌套，模拟继承
*/
func main() {

	var instance Point
	instance.name = "张三"
	instance.age = 20
	instance.address = "四川成都市"
	instance.idno = "123456789"
	instance.faceImg = "https://xxx.com/xxxsd.jpg"
	log.Printf("%v", instance)

	var pt = make(map[int]*Point)

	for index := 0; index < 10; index++ {
		pt[index] = &instance
	}

	for _, v := range pt {
		log.Printf("%v", v.name)
	}

	ins := Point{
		name:    "xxx",
		age:     20,
		address: "1123",
		idno:    "2323",
		faceImg: "https://www.baidu.com",
	}
	log.Printf("%v,%v\n", ins, ins.name)

	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	log.Printf("%v,%v\n", relation, relation.child.child.name)

	ins1 := new(Point)
	ins1.name = "sss"
	ins1.address = "333"
	log.Println(ins1)

	// 多个值列表初始化结构体
	/*
		1、必须初始化结构体的所有字段
		2、每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		3、键值对与值列表的初始化形式不能混用。
	*/
	ins2 := Point{
		"zhang",
		20,
		"xxx",
		"333",
		"https://www.baidu.com",
	}
	log.Println(ins2)

	// 匿名结构,并且初始化
	ins3 := struct {
		age     int
		address string
	}{
		age:     10,
		address: "22",
	}
	log.Println(ins3)

	ins4 := &struct {
		id   int
		data string
	}{
		100, // id:100,data:"xxx"
		"xx",
	}

	printMsgType(ins4)

	var p = getPeopleById(1)

	log.Printf("%s,%s,%v", p.name, p.child.name, p.child.child)

	var m MyInt
	log.Printf("%v\n", m.IsZero())
	m = 1
	log.Printf("%v\n", m.IsZero())

	// 内嵌结构体的字段名是它的类型名
	var black BlackColor
	black.basic.B = 1
	black.basic.R = 2
	black.basic.G = 3
	black.alpha = 4
	black.Start()
	log.Printf("%v,%v\n", black.alpha, black.basic)

	// 内嵌结构体可以直接访问成员变量
	var b PurpleColor
	b.G = 1
	b.R = 2
	b.B = 3
	b.alpha = 11
	b.pink = 5
	b.End()
	log.Printf("%v,%v", b.G, b.pink)

}

// 任意类型添加方法
type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

// 方法传递匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	log.Printf("%d,%s", msg.id, msg.data)
}

// 方法返回结构体
func getPeopleById(id int) *People {
	return &People{
		name: "zhangs",
		child: &People{
			name: "xxx",
		},
	}
}

// 定义结构体
type Point struct {
	name    string
	age     int
	address string
	idno    string
	faceImg string
}

type People struct {
	name  string
	child *People
}

type Color interface {
	Start()
	End()
}

func (bc BlackColor) Start() {
	log.Println("====start")
}

func (bc BlackColor) End() {
	log.Println("====end==", bc.alpha)
}

type BasicColor struct {
	R, G, B float32
}
type BlackColor struct {
	basic BasicColor
	alpha float32
}
type PurpleColor struct {
	BasicColor
	BlackColor
	pink float32
}
