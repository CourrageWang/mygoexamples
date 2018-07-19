package main

import (
	"os"
	"log"
	"fmt"
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
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func main() {
	//createEmptyFile()
	getFileInfo()
}
