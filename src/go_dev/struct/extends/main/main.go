package main

import "fmt"

type student struct {
	name    string
	age     int
	address string
}

func (s *student) SetName(name string) {
	s.name = name
}
func (s *student) SetAge(age int) {
	s.age = age
}
func (s *student) StudentInfo() {
	fmt.Println("姓名：", s.name, "年龄", s.age)
}

type gradute struct {
	name string
	student
}

type secondarySchool struct {
	student
}

func main() {
	var s = gradute{
		name: "zhangsn",
		student: student{
			name: "xx",
		},
	}
	s.SetAge(21)
	s.StudentInfo()
}
