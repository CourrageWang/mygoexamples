package main

import (
	"fmt"
	"sync"
	"time"
)

const APP_VER = 1.0 //
type factory struct {
	Name string
	Id   int
}

func main() {
	//fmt.Println("hello world")// 注释
	//fmt.Println(test().Name)
	//var st string  ="str"
	//fmt.Println(st)
	//str(&st)
	//fmt.Println(st)
	//Mycompare()
	//syncMap
	cur:=time.Now()
	timestamp:=cur.UnixNano()/1000000
	fmt.Println(timestamp)

}
func test() factory {
	userInfo := "Test"
	f := factory{}
	f.Name = userInfo
	f.Id = 1
	return f

}

func str(string2 *string) {
	*string2 = "test"
}
func Mycompare() {
	var i = "11"
	var j = "11"
	fmt.Println(i == j)
}
func syncMap() {
	var m sync.Map
	m.Store("name", "Mr james")
	m.Store("id", 1)
	va, _ := m.Load("id")
	fmt.Println(va)

}
