package packet

import (
	"fmt"
	"log"
	"net"
	"time"

	"ids/packet/tcpassembly"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Listen 监听网卡.
func (slf *Handle) Listen() error {
	// 获取网卡信息
	iface, err := net.InterfaceByName(slf.cardName)
	if err != nil {
		return fmt.Errorf("cardName %s not found, err: %v", slf.cardName, err)
	}
	log.Printf("cardName: %s, MTU: %d", slf.cardName, iface.MTU)

	// 打开设备监听
	handle, err := pcap.OpenLive(slf.cardName, 1024*1024, slf.promisc, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("openLive %s err: %v", slf.cardName, err)
	}
	defer handle.Close()

	// 设置过滤器
	if err := handle.SetBPFFilter(slf.bpf); err != nil {
		return fmt.Errorf("set bpf filter: %v", err)
	}

	go slf.EventHandle()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	streamFactory := NewHTTPStreamFactory(slf.eventCh)
	pool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(pool)

	ticker := time.Tick(time.Minute)
	var lastPacketTimestamp time.Time

	for {
		select {
		case <-slf.ctx.Done():
			return nil
		case packet := <-packetSource.Packets():
			netLayer := packet.NetworkLayer()
			if netLayer == nil {
				continue
			}
			transLayer := packet.TransportLayer()
			if transLayer == nil {
				continue
			}

			tcp, ok := transLayer.(*layers.TCP)
			if ok {
				assembler.AssembleWithTimestamp(
					netLayer.NetworkFlow(),
					tcp,
					packet.Metadata().CaptureInfo.Timestamp, packet)

				lastPacketTimestamp = packet.Metadata().CaptureInfo.Timestamp
			}

		case <-ticker:
			assembler.FlushOlderThan(lastPacketTimestamp.Add(slf.flushTime))
		}
	}
}

func (slf *Handle) CaptureExp80() error {
	// 获取网卡信息
	iface, err := net.InterfaceByName(slf.cardName)
	if err != nil {
		return fmt.Errorf("cardName %s not found, err: %v", slf.cardName, err)
	}
	log.Printf("cardName: %s, MTU: %d", slf.cardName, iface.MTU)

	// 打开设备监听
	handle, err := pcap.OpenLive(slf.cardName, 1024*1024, slf.promisc, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("openLive %s err: %v", slf.cardName, err)
	}
	defer handle.Close()

	// 设置过滤器
	if err := handle.SetBPFFilter(slf.bpfU); err != nil {
		return fmt.Errorf("set bpfU filter: %v", err)
	}

	go slf.EventHandleU()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		select {
		case <-slf.ctx.Done():
			return nil
		case packet := <-packetSource.Packets():
			doc := tcpassembly.Sort(packet)

			if doc != nil {
				slf.eventUCh <- EventU{
					Doc: doc,
				}
			}

		}

	}
}
