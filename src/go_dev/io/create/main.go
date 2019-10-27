package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件操作
func main() {
	createFile()
	//ReadFile()
}

type User struct {
	Name string `json:"name"`
}

func createFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("创建文件失败")
	}
	defer file.Close()
	fmt.Println("创建完成")
}

// 读取文件
func ReadFile() {
	getgid := os.Getgid()
	fmt.Println(getgid)
	dir, _ := os.Getwd()
	fmt.Printf("%v", dir)
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
