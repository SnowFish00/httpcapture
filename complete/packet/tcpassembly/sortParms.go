package tcpassembly

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// 获取可打印的ASCII字符
func getPrintableASCIIB(data []byte) string {
	printable := make([]byte, len(data))
	for i, b := range data {
		if b >= 32 && b <= 126 {
			printable[i] = b
		} else {
			printable[i] = '.'
		}
	}
	return string(printable)
}

// 处理以太网数据包
func handleEthernetPacket(packet gopacket.Packet, doc map[string]interface{}) map[string]interface{} {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	EthernetPacket, _ := ethernetLayer.(*layers.Ethernet)

	doc["SrcMAC"] = EthernetPacket.SrcMAC.String()
	doc["DstMAC"] = EthernetPacket.SrcMAC.String()
	doc["TimeStamp"] = packet.Metadata().Timestamp
	doc["Detial"] = fmt.Sprintf("%s", packet.Dump())

	return doc

}

// 处理IPv4数据包
func handleIPv4Packet(packet gopacket.Packet, doc map[string]interface{}) map[string]interface{} {
	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	IPV4Packet, _ := ipv4Layer.(*layers.IPv4)

	doc["SrcIP"] = IPV4Packet.SrcIP.String()
	doc["DstIP"] = IPV4Packet.DstIP.String()
	doc["protocol"] = IPV4Packet.Protocol.String()

	return doc

}

// 处理IPv6数据包
func handleIPv6Packet(packet gopacket.Packet, doc map[string]interface{}) map[string]interface{} {
	ipv6Layer := packet.Layer(layers.LayerTypeIPv6)
	IPV6Packet, _ := ipv6Layer.(*layers.IPv6)

	doc["SrcIP"] = IPV6Packet.SrcIP.String()
	doc["DstIP"] = IPV6Packet.DstIP.String()
	doc["NextHeader"] = IPV6Packet.NextHeader.String()

	return doc
}

// 处理TCP数据包
func handleTCPPacket(packet gopacket.Packet, doc map[string]interface{}) map[string]interface{} {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	TCPPacket, _ := tcpLayer.(*layers.TCP)

	doc["index"] = "TCP"
	doc["SrcPort"] = TCPPacket.SrcPort.String()
	doc["DstPort"] = TCPPacket.DstPort.String()
	doc["PayLoad"] = getPrintableASCIIB(TCPPacket.Payload)
	return doc

}

// 处理UDP数据包
func handleUDPPacket(packet gopacket.Packet, doc map[string]interface{}) map[string]interface{} {
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	UDPPacket, _ := udpLayer.(*layers.UDP)

	doc["index"] = "UDP"
	doc["SrcPort"] = UDPPacket.SrcPort.String()
	doc["DstPort"] = UDPPacket.DstPort.String()
	doc["PayLoad"] = getPrintableASCII(UDPPacket.Payload)
	return doc

}

func Sort(packet gopacket.Packet) map[string]interface{} {

	doc := make(map[string]interface{})
	if packet.Layer(layers.LayerTypeEthernet) != nil {
		doc = handleEthernetPacket(packet, doc)
		if packet.Layer(layers.LayerTypeIPv4) != nil {
			doc = handleIPv4Packet(packet, doc)
			if packet.Layer(layers.LayerTypeTCP) != nil {
				doc = handleTCPPacket(packet, doc)
			}
			if packet.Layer(layers.LayerTypeUDP) != nil {
				doc = handleUDPPacket(packet, doc)
			}
		}
		if packet.Layer(layers.LayerTypeIPv6) != nil {
			doc = handleIPv6Packet(packet, doc)
			if packet.Layer(layers.LayerTypeTCP) != nil {
				doc = handleTCPPacket(packet, doc)
			}
			if packet.Layer(layers.LayerTypeUDP) != nil {
				doc = handleUDPPacket(packet, doc)
			}
		}
	}

	return doc

}
