# VRC Camera Dolly Import

VRChatのカメラドリーのパスをOSC経由でインポートするためのツールです。

## 概要

このツールは、VRChatで出力したカメラドリーのJSONファイルを読み込み、OSC経由でVRChatに送信してカメラパスを再インポートできます。

## 必要要件

- Go言語の実行環境
- VRChat

## インストール

```bash
git clone https://github.com/yoshiken/vrc-camera-dolly-import
cd vrc-camera-dolly-import
go mod download
```

## 使用方法

1. VRChatでカメラドリーのパスをJSONファイルとして出力します
2. 以下のコマンドを実行してカメラパスをインポートします：

```bash
go run main.go -file "カメラパスのJSONファイルへのパス" [-ip IPアドレス] [-port ポート番号]
```

### オプション

- `-file`: カメラパスのJSONファイルパス（必須）
- `-ip`: VRChatが動作しているPCのIPアドレス（デフォルト: 127.0.0.1）
- `-port`: OSCのポート番号（デフォルト: 9000）

### 使用例

基本的な使用方法：
```bash
go run main.go -file ./camera.json
```

IPアドレスを指定：
```bash
go run main.go -file ./camera.json -ip 192.168.0.14
```

すべてのオプションを指定：
```bash
go run main.go -file ./camera.json -ip 192.168.0.14 -port 9000
```

## 注意事項

- VRChatでOSCが有効になっていることを確認してください
- ファイアウォールの設定でOSCの通信が許可されていることを確認してください
- jsonファイルが長いとうまくimportできないことがあります
- 付属のcamera.jsonはexampleです
