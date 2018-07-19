package main

import (
	"net"
	"fmt"
	"strings"
)
//http://zuozuohao.github.io/2016/06/16/Practical-Persistence-in-Go-Organising-Database-Access/
// 封装派网信息 方便业务扩展
type PannabitEvent struct {
	eventType   string      //事件类型
	detailEvent interface{} //详细信息
}
// 认证事件详细信息
type DetailusrAuth struct {
	loginTime  string
	sourceAddr string
	account    string
	macAddr    string
}
/**
  udp 服务端
 */
func main() {
	listenUdp()
}
func listenUdp() {
	// 指定端口号
	laddr := net.ParseIP("0.0.0.0")
	addr := &net.UDPAddr{IP: laddr, Port: 9527}
	conn, err := net.ListenUDP("udp4", addr,
		//	&net.UDPAddr{
		//	IP:   net.IPv4(127, 0, 0, 1),
		//	Port: 8080,
		//}

	)
	if err != nil {
		fmt.Println("listen err ...",err)
		return
	}
	defer conn.Close()
	// 监听数据
	for {
		data := make([]byte, 2*1024)
		size, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Read err ...")
			return
		}
		fmt.Println("remoteAddr:", remoteAddr)
		arrData := strings.Split(string(data[:size]), " ")
		HandleMessage(arrData)
	}

}

// 消息处理
func HandleMessage(data []string) {
	if ok := len(data); ok > 0 {
		action := data[0]
		switch action {
		case "usrauth": //认证事件
			handleAuth(data)
		default:

		}
	}
}

// 处理认证事件
func handleAuth(parms [] string) {
	detail := DetailusrAuth{
		loginTime:  parms[1],
		sourceAddr: parms[2],
		account:    parms[3],
		macAddr:    parms[4],
	}
	event := &PannabitEvent{
		eventType:   parms[0],
		detailEvent: detail,
	}
	fmt.Println(event)
	fmt.Println(event.detailEvent.(DetailusrAuth)) //强制转换为DetailusrAuth.获取到数据，做后续处理
	// 具体业务处理。。。。

}
