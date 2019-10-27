package main

import (
	"fmt"
	"sort"
)

// 定义英雄结构体
type Hero struct {
	Name   string
	Age    int
	Weight int
}

// 自定义个英雄切片类型
type HeroSlice []Hero

// 切片类型实现排序
func (h HeroSlice) Len() int {
	return len(h)
}
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age > h[j].Age
}

func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {

	// 声明自定义切片类型
	var heroSlice HeroSlice

	for i := 0; i < 10; i++ {
		var h = Hero{
			Name:   fmt.Sprintf("xx%d", i),
			Age:    i * 10,
			Weight: 20 * i * 1,
		}
		// 循环添加到切片中
		heroSlice = append(heroSlice, h)
	}

	for _, v := range heroSlice {
		fmt.Println("value", v)
	}
	// 排序
	sort.Sort(heroSlice)
	for i := 0; i < len(heroSlice); i++ {
		fmt.Println(heroSlice[i])
	}
}
