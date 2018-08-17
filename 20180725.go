package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func file(str string) {
	file, err := os.OpenFile("bufio2.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buferedWrite := bufio.NewWriterSize(file, 200)
	_, er := buferedWrite.WriteString("这是测试")
	if er != nil {
		log.Fatal(er)
	}

	//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
	buf := []byte(" write by bytes\n")
	if _, err := buferedWrite.Write(buf); err == nil {
		fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.")
		if err := buferedWrite.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("Successful flush the buffer data to file ")
	}

	// 写字符串到buffer
	_, bserr := buferedWrite.WriteString("hello this is string ..")
	if bserr != nil {
		log.Fatal(bserr)
	}

	// 写内存buffer到磁盘
	buferedWrite.Flush()

	// 丢弃还没有flush的缓存内容
	buferedWrite.Reset(buferedWrite)
}

func main() {
	file("aaa")
	file("bbbb")

}
