package main

import (
	"fmt"
	"net"
)

//http://zuozuohao.github.io/2016/06/16/Practical-Persistence-in-Go-Organising-Database-Access/
func main() {
	fmt.Println("Client")
	localip := net.ParseIP("192.168.0.106")
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
	//str := "<PNB0>mobileagt 1532053847 192.168.0.107 oppo R11"
	str :="<PNB0>ipcookie3 1533191796 e0-cb-4e-41-66-bf 192.168.0.110 0 0.0.0.0 0 1 0 0 0"
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
