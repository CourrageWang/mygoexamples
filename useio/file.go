package main

import (
	"os"
	"log"
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
)

var (
	newFile  *os.File
	err      error
	fileInfo os.FileInfo
)

//创建一个空文件
func createEmptyFile() {
	newFile, err = os.Create("test.txt")
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)

}

//  文件信息
func getFileInfo() {
	fileInfo, err = os.Stat("test_new.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

// 文件重命名
func rNameFile() {
	oldname := "test.txt"
	newname := "test_new.txt"
	err := os.Rename(oldname, newname)
	if err != nil {
		log.Fatal(err)
	}
}

// 删除文件
func delFile() {
	err := os.Remove("PA测试.md")
	if err != nil {
		log.Fatal(err)
	}
}

// 打开文件
func openFile() {
	// 以只读的方式打开文件
	file, err := os.Open("test_new.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	// 打开时带有权限
	file2, err2 := os.OpenFile("test_new.txt", os.O_APPEND, 0666)
	defer file2.Close()
	if err2 != nil {
		log.Fatal(err)
	}
	// 下面的属性可以单独使用，也可以组合使用。
	// 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	// os.O_CREATE|os.O_APPEND
	// 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
}

// 检查文件是否存在
func checkFile() {
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exit ....")
		}
	}
}

// -----------------文件操作--------------
// 复制文件
func copyfile() {
	//打开原始文件
	originalFile, err := os.Open("test_new.txt")
	defer originalFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 创建新的文件
	newfile, nerr := os.Create("new_tets.txt")
	if nerr != nil {
		log.Fatal(nerr)
	}
	defer newfile.Close()
	// 开始复制文件
	bytesWritten, werr := io.Copy(newfile, originalFile)
	if werr != nil {
		log.Fatal(werr)
	}
	log.Printf("Copied %d bytes", bytesWritten)
	// 将文件内容flush到磁盘
	serr := newfile.Sync()
	if serr != nil {
		log.Fatal(serr)
	}
}

// 写文件
func writeFile() {
	file, err := os.OpenFile("new_tets.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte("I am the auth!")
	bytenWritten, werr := file.Write(byteSlice)
	if werr != nil {
		log.Fatal(werr)
	}
	log.Fatalf("Worte %d bytes", bytenWritten)

}

//快速写文件
func quickwriteFile() {
	err := ioutil.WriteFile("quickwritefile.txt", []byte("hello"), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

// 使用缓存写文件
func writefileUsebuf() {
	file, err := os.OpenFile("bufio.txt", os.O_WRONLY, 0666)
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
	// 检查混存中的字节数
	hasBufferedSize := buferedWrite.Buffered()
	log.Printf("bytes buffered %d", hasBufferedSize)
	// 还用多少字节可用
	notBuffered := buferedWrite.Available()
	log.Printf("Auailable buffer %d", notBuffered)

	// 写内存buffer到磁盘
	buferedWrite.Flush()

	// 丢弃还没有flush的缓存内容
	buferedWrite.Reset(buferedWrite)
	bytesAvailable := buferedWrite.Available()
	log.Printf("Auailable buffer %d", bytesAvailable)
}

// 读取最多N个字节
func mostRead() {
	file, err := os.Open("new_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := make([] byte, 16)
	_, err2 := file.Read(byteSlice)
	if err2 != nil {
		log.Fatal(err2)
	}
	//log.Fatalf("has readed %d", len)
	log.Fatalf("data is :%s", byteSlice)
}

// 读取全部字节
func readAll() {
	file, err := os.Open("new_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Printf("data is %s", data)
}

//读取文件到内存
func ReadToMemory() {

	// 读取文件到byte slice中
	data, err := ioutil.ReadFile("new_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatalf("the data is %s", data)
}

// 使用缓存读
func readUseBuf() {
	file, err2 := os.Open("new_test.txt")
	if err2 != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buferedReader := bufio.NewReader(file)
	//  得到字节 ，当前指针不变
	byteslice := make([]byte, 5)
	byteslice, err = buferedReader.Peek(5) // peek返回输入流的下n个字节。而不会移动读取的位置。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("peekde at 5 bytes %s\n ", byteslice)

	// 读取同时移动指针
	_, rerr := buferedReader.Read(byteslice)
	if rerr != nil {
		log.Fatal(rerr)
	}
	fmt.Printf("read data is %s\n", byteslice)

	// 读取到分割符
	data, byerr := buferedReader.ReadBytes('\n')
	if byerr != nil {
		log.Fatal(byerr)
	}
	fmt.Printf("the data is:%s", data)

}
func main() {
	//createEmptyFile()
	//getFileInfo()
	//rNameFile()
	//delFile()
	//checkFile()
	//copyfile()
	//writeFile()
	//quickwriteFile()
	//writefileUsebuf()
	//mostRead()
	//readAll()
	//ReadToMemory()
	//readUseBuf()
	writefileUsebuf()
}
