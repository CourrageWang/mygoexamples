package main

import (
	"fmt"
	"os"
)

// 文件基本使用

func openFile() {
	file, err := os.Open("/Users/yqwang/Workspace/gopath/src/github.com/mygoexamples/Test.txt")
	if err != nil {
		fmt.Print("file open error", err)
	}

	fmt.Println("file=%V", file)

	err = file.Close()
	if err !=nil {
		fmt.Println("close file error" ,err)
	}
}

func main() {
	//openFile()
	fmt.Println(111)
}
