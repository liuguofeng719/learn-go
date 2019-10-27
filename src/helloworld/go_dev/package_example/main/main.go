package main

import (
	"fmt"
	"HelloWord/go_dev/package_example/calc"
)

var name string ="123"
var Name string ="345"

func main() {
	sum := calc.Add(10,20)
	sub := calc.Sub(20,10)
	fmt.Println(sum,sub);
}