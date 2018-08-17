package main

import (
	"fmt"
)

var (
	flags bool
	strs  string
)

func foo(ch chan strisng) {
	flags = true
	strs = "setup complete!"
	ch <- "I'm complete." //foo():我的任务完成了，发个消息给你~
}

func main() {
	ch := make(chan string)
	go foo(ch)
	<-ch //main():OK，收到你的消息了~
	for !flags {
	}
	fmt.Println(strs)
}

