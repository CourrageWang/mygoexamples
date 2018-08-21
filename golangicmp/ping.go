package main

import (
	"net"
)

func main() {
	laddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")}
	raddr := net.IPAddr{IP: net.ParseIP("192.168.0.122")}
	net.DialIP("ip4:icmp", &laddr, raddr)
}
