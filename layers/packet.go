package layers

import "fmt"

// Packet は各層間で受け渡しを行う共通のインターフェースです。各層のデータを構造体で表現します。
type Packet interface {
	Show(layerName string)
}

// ApplicationPacket はApplication Layerのデータを表す構造体です。
type ApplicationPacket struct {
	Payload string
}

func (p ApplicationPacket) Show(layerName string) {
	fmt.Printf("[%s]\n", layerName)
	fmt.Printf("---Application Layer Fields---\n")
	fmt.Printf("Payload: %s\n", p.Payload)
	fmt.Println()
}

// TransportPacket はTransport Layerのデータを表す構造体です。
type TransportPacket struct {
	SrcPort int
	DstPort int
	ApplicationPacket ApplicationPacket
}

func (p TransportPacket) Show(layerName string) {
	fmt.Printf("[%s]\n", layerName)
	fmt.Printf("---Application Layer Fields---\n")
	fmt.Printf("Payload: ***hidden***\n")
	fmt.Printf("---Transport Layer Fields---\n")
	fmt.Printf("SrcPort: %d\n", p.SrcPort)
	fmt.Printf("DstPort: %d\n", p.DstPort)
	fmt.Println()
}

// InternetPacket はInternet Layerのデータを表す構造体です。
type InternetPacket struct {
	SrcIP   string
	DstIP   string
	TransportPacket TransportPacket
}

func (p InternetPacket) Show(layerName string) {
	fmt.Printf("[%s]\n", layerName)
	fmt.Printf("---Application Layer Fields---\n")
	fmt.Printf("Payload: ***hidden***\n")
	fmt.Printf("---Transport Layer Fields---\n")
	fmt.Printf("SrcPort: ***hidden***\n")
	fmt.Printf("DstPort: ***hidden***\n")
	fmt.Printf("---Internet Layer Fields---\n")
	fmt.Printf("SrcIP: %s\n", p.SrcIP)
	fmt.Printf("DstIP: %s\n", p.DstIP)
	fmt.Println()
}

// NetworkInterfacePacket はNetwork Interface Layerのデータを表す構造体です。
type NetworkInterfacePacket struct {
	Preamble  string
	DstMAC    string
	SrcMAC    string
	EtherType string
	InternetPacket InternetPacket
	FCS       string
}

func (p NetworkInterfacePacket) Show(layerName string) {
	preambleShort := p.Preamble
	if len(preambleShort) > 10 {
		preambleShort = preambleShort[:10] + "..."
	}
	fmt.Printf("[%s]\n", layerName)
	fmt.Printf("---Application Layer Fields---\n")
	fmt.Printf("Payload: ***hidden***\n")
	fmt.Printf("---Transport Layer Fields---\n")
	fmt.Printf("SrcPort: ***hidden***\n")
	fmt.Printf("DstPort: ***hidden***\n")
	fmt.Printf("---Internet Layer Fields---\n")
	fmt.Printf("SrcIP: ***hidden***\n")
	fmt.Printf("DstIP: ***hidden***\n")
	fmt.Printf("---Network Interface Layer Fields---\n")
	fmt.Printf("Preamble: %s\n", preambleShort)
	fmt.Printf("DstMAC: %s\n", p.DstMAC)
	fmt.Printf("SrcMAC: %s\n", p.SrcMAC)
	fmt.Printf("EtherType: %s\n", p.EtherType)
	fmt.Printf("FCS: %s\n", p.FCS)
	fmt.Println()
}
