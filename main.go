package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	// デフォルト値の設定
	defaultIP := "127.0.0.1"
	defaultPort := 9000

	// コマンドライン引数の定義
	filePath := flag.String("file", "", "カメラパスのJSONファイルパス")
	ip := flag.String("ip", defaultIP, "VRChatが動作しているPCのIPアドレス")
	port := flag.Int("port", defaultPort, "OSCのポート番号")

	// ヘルプメッセージのカスタマイズ
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "使用方法: %s [options] \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n例:\n")
		fmt.Fprintf(os.Stderr, "  %s -file camera.json\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -file camera.json -ip 192.168.0.14\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -file camera.json -ip 192.168.0.14 -port 9000\n", os.Args[0])
	}

	flag.Parse()

	// ファイルパスが指定されていない場合はヘルプを表示
	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// JSONファイルを読み込む
	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// OSCクライアントを初期化
	client := osc.NewClient(*ip, *port)

	// JSONデータを文字列として送信
	msg := osc.NewMessage("/dolly/Import", string(data))
	err = client.Send(msg)
	if err != nil {
		fmt.Printf("Error sending OSC message: %v\n", err)
		return
	}

	fmt.Printf("Sending camera path to %s:%d\n", *ip, *port)
	fmt.Println("Camera path import completed!")
}
