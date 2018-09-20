package main

/**
    ------------------ DHCP service....------------------
http://blog.51cto.com/tonyguo/163475
 */

import (
	dhcp "github.com/krolaw/dhcp4"

	"log"
	"math/rand"
	"net"
	"time"
	"fmt"
)

// 使用DHCP与单个网络接口设备的示例
func main() {
	log.Print("正在开启dhcp服务器......")
	serverIP := net.IP{192, 168, 0, 125} //serverIp 172.30.0.1 //192.168.0.101

	handler := &DHCPHandler{
		ip:            serverIP,
		leaseDuration: 5 * time.Second,
		start:         net.IP{192, 168, 0, 120},
		leaseRange:    3,
		leases:        make(map[int]lease, 10),
		options: dhcp.Options{
			dhcp.OptionSubnetMask:       []byte{255, 255, 255, 0},       //子网掩码
			dhcp.OptionRouter:           []byte(net.IP{192, 168, 0, 1}), // 假设服务是你的路由器
			dhcp.OptionDomainNameServer: []byte(net.IP{61, 134, 1, 4}),  // 假设服务器是你的dns服务器
		},
	}
	log.Fatal(dhcp.ListenAndServe(handler))
	//log.Fatal(dhcp.Serve(dhcp.NewUDP4BoundListener("eth0",":67"), handler)) // Select interface on multi interface device - just linux for now
	// log.Fatal(dhcp.Serve(dhcp.NewUDP4FilterListener("en0",":67"), handler)) // Work around for other OSes
}

type lease struct {
	nic    string    // Client's CHAddr
	expiry time.Time // 租约到期时
}

type DHCPHandler struct {
	ip            net.IP        // 要使用的服务器的Ip
	options       dhcp.Options  // 发送到DHCP客户端的选项
	start         net.IP        // 分配IP范围的开始
	leaseRange    int           // 要分发的IP数量（从头开始）
	leaseDuration time.Duration // 租赁期
	leases        map[int]lease // 映射以跟踪租约
}

func (h *DHCPHandler) ServeDHCP(p dhcp.Packet, msgType dhcp.MessageType, options dhcp.Options) (d dhcp.Packet) {
	switch msgType {

	case dhcp.Discover:
		free, nic := -1, p.CHAddr().String()
		for i, v := range h.leases { // Find previous lease
			if v.nic == nic {
				free = i
				log.Println("I am at here1 ", free)
				goto reply
			}
		}
		if free = h.freeLease(); free == -1 {
			log.Println("I am at here2 ", free)
			return
		}
	reply:
		return dhcp.ReplyPacket(p, dhcp.Offer, h.ip, dhcp.IPAdd(h.start, free), h.leaseDuration,
			h.options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]))

	case dhcp.Request:
		if server, ok := options[dhcp.OptionServerIdentifier]; ok && !net.IP(server).Equal(h.ip) {
			return nil //
		}
		reqIP := net.IP(options[dhcp.OptionRequestedIPAddress])
		if reqIP == nil {
			reqIP = net.IP(p.CIAddr())
		}

		if len(reqIP) == 4 && !reqIP.Equal(net.IPv4zero) {
			if leaseNum := dhcp.IPRange(h.start, reqIP) - 1; leaseNum >= 0 && leaseNum < h.leaseRange {
				if l, exists := h.leases[leaseNum]; !exists || l.nic == p.CHAddr().String() {
					h.leases[leaseNum] = lease{nic: p.CHAddr().String(), expiry: time.Now().Add(h.leaseDuration)}
					return dhcp.ReplyPacket(p, dhcp.ACK, h.ip, reqIP, h.leaseDuration,
						h.options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]))
				}
			}
		}
		return dhcp.ReplyPacket(p, dhcp.NAK, h.ip, nil, 0, nil)

	case dhcp.Release, dhcp.Decline:
		nic := p.CHAddr().String()
		for i, v := range h.leases {
			if v.nic == nic {
				fmt.Println("ip冲突 。。。..." , h.leases )
				delete(h.leases, i)
				break
			}
		}
	}
	return nil
}

func (h *DHCPHandler) freeLease() int {
	now := time.Now()
	b := rand.Intn(h.leaseRange) // Try random first
	for _, v := range [][]int{[]int{b, h.leaseRange}, []int{0, b}} {
		for i := v[0]; i < v[1]; i++ {
			if l, ok := h.leases[i]; !ok || l.expiry.Before(now) {
				log.Println("I am at here3 ", )
				return i
			}
		}
	}
	return -1
}
