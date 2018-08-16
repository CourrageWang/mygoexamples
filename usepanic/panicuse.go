package main

import "fmt"

func main() {
	co(0)
	co2()
	co(1)
}

/**
 函数中如果书写了panic语句，会终止后面的语句。
  panic 抛出异常， recover 捕获异常。
 */
func co(i int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	fmt.Println(12 / i)
	fmt.Print("nomal  run ")

}
func co2() {
	fmt.Println(" this at  co2")
}
