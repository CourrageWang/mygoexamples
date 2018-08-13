package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
    golang 获取随机数
 */
func main() {
	fmt.Println(rand.Intn(100)) // 随机数函数如果不设置种子函数的， 每次产生的随机数是一样的，计算机的随机是一种伪随机
	rand.Seed(time.Now().Unix()) //用时间戳作为种子函数
	fmt.Println(rand.Intn(100))
}
