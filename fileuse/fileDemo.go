package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const fileName = "/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/Test.txt"
const fileName2 = "/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/Test2.txt"

// 基本的打开文件的方式
func openFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("open file error:", err)
	}
	fmt.Println("file=%V", file)

	// 关闭文件
	defer file.Close()
}

// 待缓冲的文件并显示文件。
func openFileWithBuffer() {
	file, err := os.Open("/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/Test.txt")
	if err != nil {
		fmt.Print("open file error:", err)
	}

	// 关闭文件
	defer file.Close()

	// 创建Reader 带缓冲区 [默认4096个字节]
	reader := bufio.NewReader(file)

	for ; ; {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件的末尾
			break
		}
		fmt.Print(str)
	}

}

// 一次性的将所有文件全部读完【适合较小的文件】
func readFileWithOneTime() {
	content, err := ioutil.ReadFile(fileName) // 文件的打开与关闭被封装到ReadFile中了

	if err != nil {
		fmt.Print("read file error", err)
	}

	fmt.Print("%V", content) //[]byte
	fmt.Print("%V", string(content))

	//因为没有显示的打开文件，因此不需要显示的关闭。
}

//------------------------------------------
// 写入文件并带有缓冲区
func writeFileWithBuffer() {
	file, err := os.OpenFile(fileName2, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open file err ", err)
		return
	}

	defer file.Close()
	str := "hello Gardon"

	// 写入时 使用带缓冲的 *Write
	write := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		write.WriteString(str + "\r\n")
	}

	//write.WriteString 先写在buffer中去的，因此需要调用flush 将缓存的文件写入磁盘中去。
	write.Flush()
}

//追加文件并且覆盖掉原先的文件
func fugaiFileWithbuffer() {
	file, err := os.OpenFile(fileName2, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open file err ", err)
		return
	}
	defer file.Close()

	str := "hello Golang\r\n"

	write := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		write.WriteString(str)
	}

	write.Flush()
}

//打开一个文件并追加内容 [例如日志文件]
func appendFile() {
	file, err := os.OpenFile(fileName2, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err ", err)
		return
	}
	defer file.Close()

	str := "hello English\r\n"

	write := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		write.WriteString(str)
	}

	write.Flush()
}

// 将一个文件的内容复制到另一个文件中去
func writeFileToAnother() {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error", err)
		return
	}

	err = ioutil.WriteFile(fileName2, data, 0666)
	if err != nil {
		fmt.Println("Write file error", err)
	}
}

// 构建函数，接收两个路径【原路径、目标路径】完成文件拷贝
// 该部分可以拷贝大文件 因为其是拷贝一部分写一部分。节省空间。

func copyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file error", err)
	}
	defer srcFile.Close()

	// 通过src获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开目标路径
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open file error ", err)
	}

	defer dstFile.Close()
	// 获取writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {

	dstPath:="/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/b.gif"
	srcPath:="/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/a.gif"

	_,err:=copyFile(dstPath,srcPath)

	if err!=nil {
		fmt.Println("copy error",err)
		return
	}
	fmt.Println("copy ok!")
}