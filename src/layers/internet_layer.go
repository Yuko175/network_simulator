package layers

type InternetLayer struct {
	srcIP string
	dstIP string
}

func NewInternetLayer(src, dst string) *InternetLayer {
	return &InternetLayer{srcIP: src, dstIP: dst}
}

func (internetLayer *InternetLayer) Name() string { return "InternetLayer" }

func (internetLayer *InternetLayer) Encapsulate(packet Packet) Packet {
	internetPacket, ok := packet.(InternetPacket)
	if !ok {
		return packet
	}
	internetPacket.SrcIP = internetLayer.srcIP
	internetPacket.DstIP = internetLayer.dstIP
	internetPacket.Show(internetLayer.Name())

	return NetworkInterfacePacket{
		InternetPacket: internetPacket,
	}
}

func (internetLayer *InternetLayer) Decapsulate(packet Packet) Packet {
	packet.Show(internetLayer.Name())
	return TransportPacket{
		ApplicationPacket: packet.(InternetPacket).TransportPacket.ApplicationPacket,
		SrcPort:          packet.(InternetPacket).TransportPacket.SrcPort,
		DstPort:          packet.(InternetPacket).TransportPacket.DstPort,
	}
}