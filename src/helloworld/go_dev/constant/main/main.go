package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LType int

const (
	NONE LType = iota
	JAVA
	C
	NET
)

func (c LType) String() string {
	switch c {
	case NONE:
		return "none"
	case JAVA:
		return "java"
	case NET:
		return "net"
	}
	return "N/A"
}

func main() {
	// fmt.Printf("%s %d", JAVA, JAVA)
	// test1()
	// test2()
	// testArray()
	// typeConvert()
	// sliceTest()
	// strTest()
	// mapTest()
	// stringsTest()
	// trimSpaceTest()
	// trimTest()
	// splitTest()
	// joinTest()
}

func joinTest() {
	var str = []string{"java", "c#", "c++"}
	var joinStr = strings.Join(str, ",")
	fmt.Println("join = " + joinStr)
}

/*
	分割字符串
*/
func splitTest() {
	var str = "a b c"
	// 默认使用空格分割字符串，且返回字符串切片
	sliceStr := strings.Fields(str)
	fmt.Println(sliceStr)
	for _, value := range sliceStr {
		fmt.Println(value)
	}

	// 指定分割字符
	var str1 = "a,b,c"
	var strSplit []string = strings.Split(str1, ",")
	fmt.Println(strSplit)
	for _, value := range strSplit {
		fmt.Println(value)
	}
}

/**
trim("oldstr","需要清除的字符串")
eg:
	trim("!!!12345LL","!L") 12345
*/
func trimTest() {
	var trimStrm = "!!!12345LL"
	//12345
	newStr := strings.Trim(trimStrm, "!L")
	fmt.Printf("%s\n", newStr)

	//12345xxx
	trimLeftStr := strings.TrimLeft("xxx12345xxx", "x")
	fmt.Printf("%s\n", trimLeftStr)

	//xxx12345
	trimRightStr := strings.TrimRight("xxx12345xxx", "x")
	fmt.Printf("%s\n", trimRightStr)

	//这里的函数被循环调用里面的rune,unicode
	trimLeftFunc := strings.TrimLeftFunc("xxx12345xxx", func(r rune) bool {
		fmt.Println(r)
		return true
	})
	fmt.Printf("%s\n", trimLeftFunc)
}

/**
* strings 工具类
 */
func stringsTest() {
	// var strUrl string
	// fmt.Scanf("%s", &strUrl)
	// fmt.Println(isPrefix(strUrl))
	// fmt.Println(isHasSuffix(strUrl))
	tsl := replace()
	fmt.Println(tsl)
}

/*
清除字符串前后空格
*/
func trimSpaceTest() {
	var strs = " ddf "
	trimStr := strings.TrimSpace(strs)
	fmt.Printf("%s", trimStr)
}

/**
replace(原始字符串,"old repalce str","new str",count)
	count 小于0 这查找全部字符串替换，
	1这替换找到到第一个
全部替换
ReplaceAll("strs","old str","new str")
*/
func replace() string {
	var strs = "xxbbddeeff"
	var newStr string
	newStr = strings.Replace(strs, "bb", "cc", 1)
	return newStr
}

// 字符串后缀是否匹配
func isHasSuffix(suffix string) bool {
	return strings.HasSuffix(suffix, ".txt")
}

// 检查前缀是否存在
func isPrefix(strUrl string) bool {
	var url = "https://"
	return strings.HasPrefix(strUrl, url)
}

/*
* 1，创建map
* 2，添加值的2种方式
* 3，删除map的值，delete(mpa,key)
* 4，清空map就是重新make 一个map
 */
func mapTest() {
	// 创建map对象key 为string value 为int
	scence := make(map[string]int)
	// map 中插入值
	scence["k1"] = 10
	scence["k2"] = 20
	scence["k3"] = 30
	// 打印map对象
	fmt.Println(scence)
	// 获取单个key的值
	fmt.Println(scence["k1"])
	// 循环遍历
	for key, value := range scence {
		fmt.Printf("%s=%d \n", key, value)
	}
	// 第二种初始化map
	maps := map[string]string{
		"a": "12",
		"b": "13",
		"c": "34",
	}
	fmt.Println(maps)

	delete(maps, "a")

	fmt.Println(maps)

}

/**
字符串 连接
	1，双引号“” 只能单行
	2，单引号·· 可以多行操作，按照原格式输出
	3, fmt.Sprint连接字符串
*/
func strTest() {
	var strs string = "222" + "bb"
	var sts string = `
		床前明光
		疑似地霜
	`
	fmt.Printf("%v", strs)
	fmt.Printf("%v", sts)
	newStrs := fmt.Sprint(strs, sts)
	fmt.Println(newStrs)
	// 通过控制台输入数据，必须通过指针的方式使用
	var b int
	var c int
	fmt.Scanf("%d %d", &b, &c)
	fmt.Printf("d=%v,c=%v", b, c)
}

/**
 * 切片类型和数组的区别就是切片不需要分配大小
 * 切片声明：
 *	切片的索引从0开始，区间取值是左闭右开取值
	len 获取切片的长度
	cap 获取切片的容量
*/
func sliceTest() {
	// 声明切片类型为整型的，相当于切片别名
	type slice1 []int
	// 使用定义的类型
	var b = slice1{1, 2, 3, 4}
	// 声明一个没有初始化的切片，切片对象为nil，len为0
	var c []int
	// len为0的切片
	var d = []int{}
	// 切片中添加数据
	d = append(d, 1)
	// 通过make声明并且初始化内存大小，声明int切片初始1个元素,元素的值为0(因为是整型切片)，切片的容量为2
	var s = make([]int, 1, 2)
	fmt.Printf("make of %v\n", s) //make of [0 0]
	// append操作，切片会自动扩容，个数小于1024则以2的倍速扩容，大于则按照大小的1/4之一来扩容
	s = append(s, 3, 2, 5, 6)
	/** 使用copy 函数必须手动扩容,d没有手动扩容，导致数据没有copy过去，始终元素还是一个
	  eg
		  copy(d, s) 目标的d个数为长度为1，s长度为5
	*/
	d = make([]int, len(s), (cap(s) * 2))
	copy(d, s)
	fmt.Printf("copy === %v \n", d)
	fmt.Printf("capcity of %v，len %v\n", cap(s), len(s))
	fmt.Printf("append of %v\n", s)
	fmt.Printf("%v\n", c == nil)
	fmt.Printf("%v\n", len(d))
	fmt.Printf("切片本身=%v\n", b)
	fmt.Printf("获取所有切片=%v\n", b[:])
	fmt.Printf("获取3后面所有切片元素=%v\n", b[3:])
	fmt.Printf("获取切片区间的值=%v\n", b[1:3])
	fmt.Printf("开始到2的位置的元素=%v\n", b[0:2])
	fmt.Printf("切片最后一个元素=%v\n", b[len(b)-1])
	fmt.Printf("清除所有切片=%v\n", b[0:0])
	// 遍历切片
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	for _, v := range d {
		fmt.Println(v)
	}
}

/**
 * go 语言必须类型转换，不能自动转换
 * eg：
 	var a int =10
	var b int32 =10
	var c = a + b //这里都是整型，也是需要类型转换的
	var c = int32(a) + b
	转换的语法： type(转换的值) eg: int(b)
	数字转换成字符串需要导入strconv 包
*/
func typeConvert() {
	var age int = 1
	var age1 int8 = 2
	var age2 int32 = 100
	var age3 bool

	fmt.Printf("%#v\n", age)
	fmt.Printf("%v\n", age1)
	fmt.Printf("%T\n", age2)
	fmt.Printf("%t\n", age3)

	sum := age + int(age1)
	fmt.Printf("%d\n", sum)

	// var t1 string = "20" 字符串转数字
	t2, _ := strconv.Atoi("20")

	// 数字转字符串
	t3 := strconv.Itoa(30)
	fmt.Println("===", t2, t3)

}

/**
 *	测试数组类型
 **/
func testArray() {
	// 使用数组别名
	type My1 [2]int
	var a = My1{1, 2}
	// 遍历数组
	for _, value := range a {
		fmt.Printf("%v=", value)
	}
	fmt.Println()
	fmt.Printf("%v\n", a)
	// 声明并且初始化数组
	var numbers = [3]int{1, 2, 3}
	// 系统自动计算素组的大小
	numb := [...]int{4, 5, 6}
	fmt.Printf("%v\n", len(numbers))
	fmt.Printf("%v", numb)
}

/**
 * 测试变量的作用域
 * 相同的变量名称但是声明的方式不一样，互不干扰
   eg: 全局变量 var b string = "go"
   func a(){
	   作用域只能在函数内，且不会覆盖全局变量，这种简写不能什么数据类型，系统自动推导
	   b:="go"
	   会覆盖全局变量
	   b = "c"
   }
*/
var b = "Go"

func test1() {
	fmt.Println("===", b)
}

func test2() {
	b := "C"
	// var b = "G++"
	fmt.Println("==", b)
}
