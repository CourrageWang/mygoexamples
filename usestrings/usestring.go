package main

import (
	"fmt"
	"strings"
)

//  strings 测试包
func main() {

	fmt.Println(strings.HasPrefix("{hello}", "{h")) //判断s字符串是否是以prefix开头
	fmt.Println(strings.HasSuffix("{hello}", "}"))  //判断s字符串是否是以suffix结尾
	fmt.Println(strings.Contains("{hello}", "he"))  //判断s字符串是否包含substr串
	fmt.Println(strings.Trim("!!Achtung!Achtung!","!"))//将s前后端所有包含custer的值都取掉
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) //去掉s前后端的空格
	fmt.Println(strings.TrimLeft("!!Gopher!!","!"))// 去掉s前短包含的cutset字符串
	s:=[]string{"apple","banana","berr"}
	fmt.Println(strings.Join(s,",")) //将字符串拼接，之间用sep隔开


}
