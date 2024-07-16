package tcpassembly

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// 获取可打印的ASCII字符
func getPrintableASCII(data []byte) string {
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

func TorUparms(packet gopacket.Packet) map[string]interface{} {
	if packet.Layer(layers.LayerTypeIPv4) != nil {
		doc := map[string]interface{}{
			"TrafficType": "ipv4",
			"SrcMAC":      "",
			"DstMAC":      "",
			"SrcIP":       "",
			"DstIP":       "",
			"SrcPort":     "",
			"DstPort":     "",
			"Payload":     "",
			"Protocol":    "",
			"Timestamp":   packet.Metadata().Timestamp,
		}
		ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
		IPV4Packet, _ := ipv4Layer.(*layers.IPv4)
		doc["SrcIP"] = IPV4Packet.SrcIP.String()
		doc["DstIP"] = IPV4Packet.DstIP.String()
		doc["Protocol"] = IPV4Packet.Protocol.String()

		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			EthernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			doc["SrcMAC"] = EthernetPacket.SrcMAC.String()
			doc["DstMAC"] = EthernetPacket.DstMAC.String()
		}

		// 如果是TCP数据包
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			TCPPacket, _ := tcpLayer.(*layers.TCP)
			doc["SrcPort"] = TCPPacket.SrcPort.String()
			doc["DstPort"] = TCPPacket.DstPort.String()
			doc["Payload"] = getPrintableASCII(TCPPacket.Payload)
		}

		// 如果是UDP数据包
		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			UDPPacket, _ := udpLayer.(*layers.UDP)
			doc["SrcPort"] = UDPPacket.SrcPort.String()
			doc["DstPort"] = UDPPacket.DstPort.String()
			doc["Payload"] = getPrintableASCII(UDPPacket.Payload)
		}

		return doc
	}

	if packet.Layer(layers.LayerTypeIPv6) != nil {
		doc := map[string]interface{}{
			"TrafficType": "ipv6",
			"SrcMAC":      "",
			"DstMAC":      "",
			"SrcIP":       "",
			"DstIP":       "",
			"SrcPort":     "",
			"DstPort":     "",
			"Payload":     "",
			"NextHeader":  "",
			"Timestamp":   packet.Metadata().Timestamp,
		}
		ipv6Layer := packet.Layer(layers.LayerTypeIPv6)
		IPV6Packet, _ := ipv6Layer.(*layers.IPv6)
		doc["SrcIP"] = IPV6Packet.SrcIP.String()
		doc["DstIP"] = IPV6Packet.DstIP.String()
		doc["NextHeader"] = IPV6Packet.NextHeader.String()

		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			EthernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			doc["SrcMAC"] = EthernetPacket.SrcMAC.String()
			doc["DstMAC"] = EthernetPacket.DstMAC.String()
		}

		// 如果是TCP数据包
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			TCPPacket, _ := tcpLayer.(*layers.TCP)
			doc["SrcPort"] = TCPPacket.SrcPort.String()
			doc["DstPort"] = TCPPacket.DstPort.String()
			doc["Payload"] = getPrintableASCII(TCPPacket.Payload)
		}

		// 如果是UDP数据包
		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			UDPPacket, _ := udpLayer.(*layers.UDP)
			doc["SrcPort"] = UDPPacket.SrcPort.String()
			doc["DstPort"] = UDPPacket.DstPort.String()
			doc["Payload"] = getPrintableASCII(UDPPacket.Payload)
		}

		return doc
	}

	return nil
}
