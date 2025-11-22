package main

import (
	"fmt"

	"github.com/nakagukihisashi/network_simulator/src/layers"
)

func main() {
	// ここでは各層の初期化と、カプセル化/非カプセル化のシーケンスを行う

	// 上位(アプリ層) -> 下位(ネットワークインターフェース層) の順序でスライスに格納
	applicationLayer := layers.NewApplicationLayer("Hello Network", "alice@example.com", "bob@example.com")
	transportLayer := layers.NewTransportLayer(1234, 80)
	internetLayer := layers.NewInternetLayer("192.168.0.1", "192.168.0.2")
	networkInterfaceLayer := layers.NewNetworkInterfaceLayer("AA:BB:CC:DD:EE:FF", "11:22:33:44:55:66", "Copper")

	layersChain := []layers.Layer{applicationLayer, transportLayer, internetLayer, networkInterfaceLayer}

	// 送信: 上位層から順にEncapsulate
	fmt.Println("Encapsulation (送信):")
	// 初期Packetを作成
	var packet layers.Packet = layers.ApplicationPacket{}
	for _, l := range layersChain {
		packet = l.Encapsulate(packet)
	}

	// 受信: 下位層から順にDecapsulate
	fmt.Println("Decapsulation (受信):")
	for i := len(layersChain) - 1; i >= 0; i-- {
		l := layersChain[i]
		packet = l.Decapsulate(packet)
	}
}
