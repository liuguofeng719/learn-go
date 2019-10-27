package main

import (
	"fmt"
	"time"
)

func main() {
	GetNow()
	GetUnix()
	GetFormat()
	StrParseDate()
	SubTime()
}

// 2个时间戳相减
func SubTime() {
	startTime := time.Now().Unix()
	fmt.Println(startTime)
	var i int64 = 1
	for {
		i++
		if i > 100000 {
			break
		}
	}
	fmt.Println("----------")
	endTime := time.Now().Unix()
	fmt.Println(endTime)
	fmt.Printf("%v", endTime-startTime)
}

// 字符串转换date
func StrParseDate() {
	const format = "2006-01-02"
	var strDate = "2019-06-18"
	t, _ := time.Parse(format, strDate)
	fmt.Printf("%v\n", t)
}

// 格式化日期
func GetFormat() {
	// 格式化必须写成这个值，如果是年2006
	const format = "2006-01-02 15:04:05"
	var now = time.Now()
	fmt.Println(now.Format(format))
}

/*
* time.Now() 获取当前时间
 */
func GetNow() {
	fmt.Printf("%v\n", time.Now())
}

/*
* 获取时间戳
 */
func GetUnix() {
	fmt.Printf("%v\n", time.Now().Unix())
}
