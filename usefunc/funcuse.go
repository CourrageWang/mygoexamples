package main

import "fmt"

func main() {
	f1("this",2.3,'h')
}

// 变参函数
func f1(args ...interface{}) {
	if lens := len(args); lens > 0 {
		for _,a := range args {
			fmt.Println(a)
		}
	}else {
		fmt.Println("args is nil ")
	}
}
// 匿名函数
