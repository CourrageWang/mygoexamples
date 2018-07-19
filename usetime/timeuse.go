package main

import (
	"fmt"
	"time"
)

// iota 使用方式
type weekday int

const (
	Sunday    weekday = iota // const出现，会让iota初始化为0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	//fmt.Println(Sunday)
	//fmt.Println(Saturday)
	usetime()
}
func usetime()  {
	tim :=time.Now().Unix()
	fmt.Println(tim)
}
