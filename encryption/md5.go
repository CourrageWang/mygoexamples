package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

/**
     golang md5 加密使用
 */

func enMd51(source string) string {
	data := []byte(source)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
func enMd52(source string) string {
	w := md5.New()
	io.WriteString(w, source)                //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

func main() {
	fmt.Println(enMd51("123456"))
	fmt.Println(enMd52("123456"))
}
