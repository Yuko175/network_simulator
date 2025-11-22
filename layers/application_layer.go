package layers

import "fmt"

type ApplicationLayer struct {
	messageBody string
	sender      string
	recipient   string
}

// NewApplicationLayer はコンストラクタです。外部から初期値を渡す場合はこの関数を使います。
func NewApplicationLayer(body, sender, recipient string) *ApplicationLayer {
	return &ApplicationLayer{messageBody: body, sender: sender, recipient: recipient}
}

func (applicationLayer *ApplicationLayer) Name() string { return "ApplicationLayer" }

// Encapsulate はアプリケーション層がデータを生成してPacketにセットします。
func (applicationLayer *ApplicationLayer) Encapsulate(packet Packet) Packet {
	payload := fmt.Sprintf("To: %s; From: %s; Body: %s", applicationLayer.recipient, applicationLayer.sender, applicationLayer.messageBody)
	applicationPacket := ApplicationPacket{
		Payload: payload,
	}

	applicationPacket.Show(applicationLayer.Name())

	return TransportPacket{
		ApplicationPacket: applicationPacket,
	}
}

// Decapsulate はアプリケーション層で最後に呼ばれ、元のメッセージをPayloadから取得します。
func (applicationLayer *ApplicationLayer) Decapsulate(packet Packet) Packet {
	packet.Show(applicationLayer.Name())
	return packet
}
