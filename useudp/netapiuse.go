package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("192.168.0.1"))
}
