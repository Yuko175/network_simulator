# Network Simulator (Go)

このプロジェクトは、TCP/IPモデルの4階層を模倣した簡易ネットワークシミュレータです。
教育目的で作られており、実際のネットワーク通信は行いません。

実行方法:

```bash
# リポジトリ直下で以下のコマンドを実行
clear && go build -o simulator src/main/main.go && ./simulator
```

構成:

- `main.go` - エントリポイント。各層の初期化とカプセル化/非カプセル化のシーケンスを実行。
- `layers/` - 各層の実装パッケージ。
  - `application_layer.go` - アプリケーション層の実装。
  - `transport_layer.go` - トランスポート層の実装。
  - `internet_layer.go` - インターネット層の実装。
  - `network_interface_layer.go` - ネットワークインターフェース層（データリンク層と物理層を統合）の実装。
  - `packet.go` - 各層のパケット構造体。
  - `layer_interface.go` - 層の共通インターフェース。
- `simulator/` - 設計関連ドキュメント。
  - `設計書.md` - プロジェクトの設計書。
  - `クラス図.pu` - PlantUML形式のクラス図。
  - `シーケンス図.pu` - PlantUML形式のシーケンス図。

注意:

- 初学者向けにコメントを多く含めています。

サンプル出力:

```text
Encapsulation (送信):
[ApplicationLayer]
---Application Layer Fields---
Payload: To: bob@example.com; From: alice@example.com; Body: Hello Network

[TransportLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: 1234
DstPort: 80

[InternetLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: ***hidden***
DstPort: ***hidden***
---Internet Layer Fields---
SrcIP: 192.168.0.1
DstIP: 192.168.0.2

[NetworkInterfaceLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: ***hidden***
DstPort: ***hidden***
---Internet Layer Fields---
SrcIP: ***hidden***
DstIP: ***hidden***
---Network Interface Layer Fields---
Preamble: 1010101010...
DstMAC: 11:22:33:44:55:66
SrcMAC: AA:BB:CC:DD:EE:FF
EtherType: 0800
FCS: 0041

Decapsulation (受信):
[NetworkInterfaceLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: ***hidden***
DstPort: ***hidden***
---Internet Layer Fields---
SrcIP: ***hidden***
DstIP: ***hidden***
---Network Interface Layer Fields---
Preamble: 1010101010...
DstMAC: 11:22:33:44:55:66
SrcMAC: AA:BB:CC:DD:EE:FF
EtherType: 0800
FCS: 0041

[InternetLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: ***hidden***
DstPort: ***hidden***
---Internet Layer Fields---
SrcIP: 192.168.0.1
DstIP: 192.168.0.2

[TransportLayer]
---Application Layer Fields---
Payload: ***hidden***
---Transport Layer Fields---
SrcPort: 1234
DstPort: 80

[ApplicationLayer]
---Application Layer Fields---
Payload: To: bob@example.com; From: alice@example.com; Body: Hello Network
```
