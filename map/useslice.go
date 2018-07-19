package main

import "fmt"

func main() {
	useSlice()
}
func useSlice()  {
	var  rune [] int32
	for _,r :=range "hello 中文"{
		rune =append(rune,r)
	}
	fmt.Printf("%q\n",rune)
}