package main

import "fmt"

/*
匿名函数
 */
func main() {
	sums := func(x, y int) (sum int) {
		sum = x + y
		return
	}(3, 4)
	fmt.Println(sums)

}
