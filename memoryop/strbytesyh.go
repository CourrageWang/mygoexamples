package main

import (
	"unsafe"
	"strings"
	"fmt"
)

/**
  字符串作为一种不可变类型，在于字节数组做转换时需要付出沉重代价，根本原因是对底层数组字节数组的复制
  这种代价会在高并发下迅速放大
*/
func str2byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	s := strings.Repeat("abc", 3)
	b := str2byte(s)
	s2 := bytes2str(b)
	fmt.Println(b, s2)
}
