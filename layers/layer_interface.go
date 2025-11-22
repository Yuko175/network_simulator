package layers

// Layer は全ての層が実装する共通インターフェースです。
// Encapsulate: 上位層から受け取ったPacketにヘッダーを付与して下位層へ渡す。
// Decapsulate: 下位層から受け取ったPacketから自分のヘッダーを取り除き、上位層へ渡す。
// Name: ログ出力用に層の名称を返す。
type Layer interface {
	Encapsulate(packet Packet) Packet
	Decapsulate(packet Packet) Packet
	Name() string
}