package main

import (
	"math/rand"
	"time"
	"fmt"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

//  随机生成指定位数的大小写字母 model 为0表示大写，1表示小写
func getRandStr(len, model int, ) (string) {
	code := 65
	if model-0 == 1 {
		code = 97
	}
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + code
		bytes[i] = byte(b)
	}
	return string(bytes)
}

//  十进制0的ASCLL码是48
func getRandInt(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(10) + 48
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(getRandStr(5, 0))
		fmt.Println(getRandInt(4))
	}

}
