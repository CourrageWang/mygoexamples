package main

import (
	"fmt"
	"reflect"
)


func testreflect1()  {
	var i  int =23
	var f  float32 =1.23
	var l [] string = []string{"a","b","c"}
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(l))
}
func main()  {
	testreflect1()
}