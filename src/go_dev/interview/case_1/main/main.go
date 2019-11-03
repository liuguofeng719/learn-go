package main

import (
	"fmt"
)

// 学生对象
type Student struct {
	name string
	age  int
	addr string
}

// 声明一个值对象来接收结构体的对象，接受的对象的地址始终是一个
func main() {
	var g *Student
	var s1 = Student{name: "zhansan1", age: 20, addr: "chengdu1"}
	var s2 = Student{name: "zhansan2", age: 21, addr: "chengdu2"}
	var s3 = Student{name: "zhansan3", age: 22, addr: "chengdu3"}
	g = &s1
	fmt.Printf("%p,%p\n", &g, &g)
	g = &s2
	fmt.Printf("%p,%p\n", &g, &g)
	g = &s3
	fmt.Printf("%p,%p\n", &g, &g)

	stusMap := make(map[string]*Student)
	stus := []Student{
		{name: "zhansan1", age: 20, addr: "chengdu1"},
		{name: "zhansan2", age: 21, addr: "chengdu2"},
		{name: "zhansan3", age: 22, addr: "chengdu3"},
	}
	// 复制正确
	for i := 0; i < len(stus); i++ {
		stu := stus[i]
		stusMap[stu.name] = &stu
	}
	// []Student
	// -----------------
	// ptr len = 3 cap = 3
	// |zhansan1 | zhangsan2 | zhangsan3 | [3]Student
	// -----------------
	for index, _ := range stus {
		stu := stus[index]
		fmt.Printf("%p\n", &stu)
		stusMap[stu.name] = &stu
	}

	for k, v := range stusMap {
		fmt.Printf("key = %v,value = %v\n", k, v)
	}

}
