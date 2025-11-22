# Network Simulator (Go)

このプロジェクトは、OSIの7階層を模倣した簡易ネットワークシミュレータです。
教育目的で作られており、実際のネットワーク通信は行いません。

実行方法:

```bash
# ビルド
cd /Users/nakagukihisashi/Desktop/network_simulator
go build -o network_simulator

# 実行
./network_simulator
```

構成:

- `main.go` - エントリポイント。
- `layers/` - 各層の実装パッケージ。

注意:

- 初学者向けにコメントを多く含めています。

サンプル出力 (一部):

```text
--- Encapsulation (送信) ---
[ApplicationLayer] -> [APP]To: bob@example.com; From: alice@example.com; Body: Hello Network
[PresentationLayer] Encapsulate: [APP]To: bob@example.com; From: alice@example.com; Body: Hello Network -> [PRES:UTF-8]...
... (省略) ...
--- Decapsulation (受信) ---
[PhysicalLayer] Decapsulate: [PHYS:Copper]... -> [LINK:AA:BB:CC:DD:EE:FF->11:22:33:44:55:66]...
最終的なApplication層のデータ: To: bob@example.com; From: alice@example.com; Body: Hello Network
```
