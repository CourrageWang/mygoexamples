package main

import (
	"fmt"
	"net"
)

//http://zuozuohao.github.io/2016/06/16/Practical-Persistence-in-Go-Organising-Database-Access/
func main() {
	fmt.Println("Client")
	localip := net.ParseIP("192.168.0.105")
	remoteip := net.ParseIP("192.168.0.195")
	lAddr := &net.UDPAddr{IP: localip, Port: 9098} //源地址
	rAddr := &net.UDPAddr{IP: remoteip, Port: 9527}
	conn, err := net.DialUDP("udp4", lAddr, rAddr) //目标地址
	if err != nil {
		fmt.Println("connect fail !", err)
		return
	}
	fmt.Println(conn.LocalAddr())
	defer conn.Close()
	str := "natip 782345890 127.0.0.1 192.168.0.105"
	senddata := []byte(str)
	_, err = conn.Write(senddata)
	if err != nil {
		fmt.Println("send fail !", err)
		return
	}

	/*data := make([]byte, 20)
	read, remoteAddr, err := conn.ReadFromUDP(data) // ReadFromUD读取一个数据包，将有效负载拷贝到data上并返回自己数和数据包来源地址。
	if err != nil {
		fmt.Println("read fail !", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)*/
}
