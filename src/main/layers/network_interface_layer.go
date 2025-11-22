package layers

import (
	"fmt"
)

// NetworkInterfaceLayer はデータリンクと物理層を統合したTCP/IPモデルのNetwork Interface Layer。

type NetworkInterfaceLayer struct {
	srcMAC string
	dstMAC string
	medium string
}

func NewNetworkInterfaceLayer(srcMAC, dstMAC, medium string) *NetworkInterfaceLayer {
	return &NetworkInterfaceLayer{srcMAC: srcMAC, dstMAC: dstMAC, medium: medium}
}

func (networkInterfaceLayer *NetworkInterfaceLayer) Name() string { return "NetworkInterfaceLayer" }

func (networkInterfaceLayer *NetworkInterfaceLayer) Encapsulate(packet Packet) Packet {
	netPacket, ok := packet.(NetworkInterfacePacket)
	if !ok {
		return packet
	}

	networkInterfacePacket := NetworkInterfacePacket{
		Preamble:  "10101010101010101010101010101010101010101010101010101011", // 7 bytes + SFD
		DstMAC:    networkInterfaceLayer.dstMAC,
		SrcMAC:    networkInterfaceLayer.srcMAC,
		EtherType: networkInterfaceLayer.getEtherType(netPacket.InternetPacket),
		InternetPacket: netPacket.InternetPacket,
		FCS:       fmt.Sprintf("%04x", len(netPacket.InternetPacket.TransportPacket.ApplicationPacket.Payload)), // Simple length as FCS
	}
	networkInterfacePacket.Show(networkInterfaceLayer.Name())
	
	return networkInterfacePacket
}

func (networkInterfaceLayer *NetworkInterfaceLayer) getEtherType(internetPacket InternetPacket) string {
	// 簡略化のため、常にIPv4のEtherTypeを返す
	return "0800"
}

func (networkInterfaceLayer *NetworkInterfaceLayer) Decapsulate(packet Packet) Packet {
	packet.Show(networkInterfaceLayer.Name())
	return InternetPacket {
		TransportPacket: packet.(NetworkInterfacePacket).InternetPacket.TransportPacket,
		SrcIP:           packet.(NetworkInterfacePacket).InternetPacket.SrcIP,
		DstIP:           packet.(NetworkInterfacePacket).InternetPacket.DstIP,
	}
}
