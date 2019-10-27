package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件操作
func main() {
	readFile()
}

// 读取文件
func readFile() {
	// 获取当前目录
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	var path string = pwd + "/src/go_dev/io/file/test.log"
	file, _ := os.Open(path)
	defer file.Close()
	fmt.Printf("%T\n", file)
	// 创建一个带缓冲的Reader ,默认缓冲是4096
	reader := bufio.NewReader(file)
	for {
		// 通过空格读取，需要注意文件内容最后必须多一个空格，不然最后一个字符不能读取
		str, err := reader.ReadString(' ')
		// str, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(str))
	}
}
