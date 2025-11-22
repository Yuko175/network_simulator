package layers

type TransportLayer struct {
	srcPort int
	dstPort int
}

func NewTransportLayer(src, dst int) *TransportLayer {
	return &TransportLayer{srcPort: src, dstPort: dst}
}

func (transportLayer *TransportLayer) Name() string { return "TransportLayer" }

func (transportLayer *TransportLayer) Encapsulate(packet Packet) Packet {
	transportPacket, ok := packet.(TransportPacket)
	if !ok {
		return packet
	}
	transportPacket.SrcPort = transportLayer.srcPort
	transportPacket.DstPort = transportLayer.dstPort
	transportPacket.Show(transportLayer.Name())

	return InternetPacket{
		TransportPacket: transportPacket,
	}
}

func (transportLayer *TransportLayer) Decapsulate(packet Packet) Packet {
	packet.Show(transportLayer.Name())
	return ApplicationPacket{
		Payload: packet.(TransportPacket).ApplicationPacket.Payload,
	}
}