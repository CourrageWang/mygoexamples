package main

import (
	"net"
	"git/imc/src/floger"
	"strings"
	"fmt"
)

//panabit 事件检测

type PanabitEvent struct {
	EventType   string      // 事件类型
	DetailEvent interface{} // 详细信息
}

//共享用户检测事件
type detailNatip struct {
	EventTime string // 检测到事件的时间
	GateWayIp string // 网关地址
	PrivateIp string // 私有IP地址
}

func PanabitListenReceive() {
	laddr := net.ParseIP("0.0.0.0")
	addr := &net.UDPAddr{IP: laddr, Port: 9527}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		floger.Error("UDP发生错误", err.Error())
		return
	}
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		size, _, err := conn.ReadFromUDP(data)
		if err != nil {
			floger.Error("UDP读取错误", err.Error())
		}
		origData := string(data[:size])
		arrData := strings.Split(origData, " ")
		floger.Debug5("获取到的数据为：", origData)

		handleNatIp(arrData)
	}
}

//消息处理
func handleMessage(data []string) {
	if ok := len(data); ok > 0 {
		action := data[0]
		switch action {
		case "natip": // 共享用户检测
			handleNatIp(data)
		default:

		}

	}

}

//共享用户事件检测事件
func handleNatIp(params []string) {
	detailNatIP := detailNatip{PrivateIp: params[4]} // 获取用户ip
	usrIP := detailNatIP.PrivateIp
	fmt.Println(usrIP)

}
