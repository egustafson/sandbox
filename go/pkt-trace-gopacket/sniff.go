package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	iface   = "lo"
	snaplen = 1600 // bytes to capture from the head of the packet
	promisc = true // promiscuous mode
	filter  = "tcp[tcpflags] & (tcp-syn|tcp-ack) == tcp-syn|tcp-ack and port 2379"
)

func main() {
	if handle, err := pcap.OpenLive(iface, snaplen, promisc, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter(filter); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			// fmt.Println(packet.Dump())
			if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)
				dest := ip.DstIP
				fmt.Printf("-> %s\n", dest.To4())
			}
		}
	}
}
