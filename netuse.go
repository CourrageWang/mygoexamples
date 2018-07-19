package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	LsiteningAccept()
}

func LsiteningAccept() {
	fmt.Println("start listening ....")
	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", 8080))
	if err != nil {
		// handle error
		fmt.Println("listen failed ...", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(" Accept failed ..", err)
			continue
		}
		go func(con net.Conn) {
			buf := make([]byte, 256)
			defer con.Close()
			n, err := con.Read(buf)
			fmt.Println()
			if err != nil {
				fmt.Println("Read filed ...", err)
				return
			}
			bufData := string(buf[0:n])
			arrData := strings.Split(bufData, " ")
			fmt.Println("bufData:", arrData)
			handlesRequest(arrData)

		}(conn)
	}
}

// 消息处理
func handlesRequest(arrData []string) {
	if length := len(arrData); length > 0 {
		fmt.Println(arrData[0])
	}

}
