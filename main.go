package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// IDトークンからヘッダー・ペイロードを入手する
func main() {
	idToken := ""

	dataArray := strings.Split(idToken, ".")
	header, payload, _ := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64デコードする(jsonに戻す)
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// payloadをbase64デコードする(jsonに戻す)
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// []byte型からString型へパース
	fmt.Println("header: ", string(headerData))
	fmt.Println("payload: ", string(payloadData))
}
