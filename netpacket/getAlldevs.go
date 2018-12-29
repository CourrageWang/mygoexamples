package main

import (
	_ "github.com/google/gopacket"
	"fmt"
	"github.com/google/gopacket/pcap"
	"time"
	"log"
	"github.com/google/gopacket"

	"github.com/google/gopacket/layers"
	"strings"
)

var (
	device       string = "br0"
	snapshot_len int32  = 1024
	promiscuous  bool   = true
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle

	ethLayer    layers.Ethernet
	ipLayer     layers.IPv4
	tcpLayer    layers.TCP
	packetCount int    = 0
	port        string = "3306"
)

func main() {
	//Decoderpacket()
	//Monitorequipment()
	//getPcap()
	Filter()

}

// 监听设备捕获信息
func Monitorequipment() {
	// 打开一个设备（device）每个数据包的最大为(snapshot_len) promiscuous 是否设置为混杂模式，以及超时时间。
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSources := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSources.Packets() {
		printPacketInfo(packet)
	}
}

//  生成 pcap格式
/*func getPcap() {
	f, _ := os.Create("/root/go/src/netpacket/test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(1024, layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", device, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}

}*/

// 解码数据包

func Decoderpacket() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&ipLayer,
			&tcpLayer,
		)
		foundLayerTypes := []gopacket.LayerType{}

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			fmt.Println("Trouble decoding layers: ", err)
		}

		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeIPv4 {
				fmt.Println("IPv4: ", ipLayer.SrcIP, "->", ipLayer.DstIP)
			}
			if layerType == layers.LayerTypeTCP {
				fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
				fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
			}
		}
	}

}

// 获取所有设备
func getAlldevices() {
	fmt.Println("find devices...")
	devices, err := pcap.FindAllDevs()
	fmt.Print(err)
	if err != nil {
		fmt.Println("FindAllDevs err.", err)
	}
	fmt.Println("---------device list---------")
	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Address:", device.Addresses)
		fmt.Println("Description:", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("_ Ip Adrress:", address.IP)
			fmt.Println("_ Netmask:", address.Netmask)
		}

	}
}

// 设置过滤器
func Filter() {
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	//var filter string = "tcp and port 3306 "
	//var filter string = " (port 80 and port 443) and not host 192.168.0.1"
	var filter string = " host 192.168.0.136 "
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal("err filter ", err)
	}
	fmt.Println("Only capturing TCP port 3306 packets.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		printPacketInfo(packet)
	}

}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println("ethernetPacket.Contents:", string(ethernetPacket.Contents))
		fmt.Println("ethernetPacket.Payload:", string(ethernetPacket.Payload))
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println("ip.Payload:", string(ip.Payload))
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println("tcp.Contents:", string(tcp.Contents))
		fmt.Println("tcp.LayerContents():", string(tcp.LayerContents()))

	}

	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
		fmt.Println("layer.LayerContents():", string(layer.LayerContents()))
	}

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")

		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {

			fmt.Println("HTTP found!")
			fmt.Println(string(applicationLayer.Payload()))

		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}
