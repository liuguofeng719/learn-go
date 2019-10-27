package main

import (
	"log"
	"sort"
)

// 排序接口 sort.Interface ，可以对任意类型排序，自定义排序需要实现这3个接口
// 系统提供对字符串StringSlice IntSlice Float64Slice 排序，默认都是升序
/** type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	// 如果i < j 升序，i > j 降序
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
// eg 实现结构体排序
func main() {
	var sliceInt = []int{4, 7, 1, 9}
	// 调用系统的整形排序
	sort.Ints(sliceInt)
	for _, v := range sliceInt {
		log.Println(v)
	}

	var strsList = MyStringList{
		"java",
		"c++",
		"go",
		"c",
		"python",
	}
	log.Println("====自定义升序排序=====")
	sort.Strings(strsList)
	for _, v := range strsList {
		log.Printf("%v\n", v)
	}
	log.Println("====自定义降序排序=====")
	// 自定义排序
	sort.Sort(strsList)
	for _, v := range strsList {
		log.Printf("%v\n", v)
	}

	// 切片结构体排序
	var hero = Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(hero)

	for _, v := range hero {
		log.Println(v)
	}

	var h = []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}
	// 稳定排序
	sort.SliceStable(h, func(i, j int) bool {
		if h[i].Kind != h[j].Kind {
			return h[i].Kind < h[j].Kind
		}
		return h[i].Name < h[j].Name
	})
	for _, v := range h {
		log.Println(v)
	}
}

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

// 定义一个英雄切片指针
type Heros []*Hero

// 实现排序接口
func (h Heros) Len() int {
	return len(h)
}
func (h Heros) Less(i, j int) bool {
	// 如果分类不一致按照分类升序
	if h[i].Kind != h[j].Kind {
		return h[i].Kind < h[j].Kind
	}
	// 分类一致通过ASCII码升序排列
	return h[i].Name < h[i].Name
}

func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 定义string类型的切片
type MyStringList []string

// string类型的切片实现排序的3个接口
// 计算排序的元素的长度
func (m MyStringList) Len() int {
	return len(m)
}

// i < j 升序，i > j 降序
func (m MyStringList) Less(i, j int) bool {
	return m[i] > m[j]
}

// 交互2个元素的位置
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
