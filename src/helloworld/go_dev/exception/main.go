package main

import (
	"errors"
	"fmt"
)

func main() {
	// div()
	success, err := readConf("conf.ini")
	if err != nil {
		panic("读取文件异常")
	} else {
		fmt.Println(success)
	}
}

func div() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("出现异常=", err)
		}
	}()

	num := 10
	num2 := 0
	n1 := num / num2
	fmt.Println(n1)
}

func readConf(fileName string) (success bool, err error) {
	if fileName == "conf.ini" {
		return true, nil
	} else {
		return false, errors.New("读取文件失败")
	}
}
