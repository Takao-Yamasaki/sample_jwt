package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

// IDトークンからヘッダー・ペイロード・署名を入手する
func main() {
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjAyYTYxZWZkMmE0NGZjMjE1MTQ4ZDRlZmZjMzRkNmE3YjJhYzI2ZjAiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiMTAyOTU5MzE0NjUyMy01MjdrMG82NzVlaGZlM2ljaHRxcW1tMDR1cHFtaTN1NC5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsImF1ZCI6IjEwMjk1OTMxNDY1MjMtNTI3azBvNjc1ZWhmZTNpY2h0cXFtbTA0dXBxbWkzdTQuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDMyNzY2MTI2MzEyMzcyMzkxMjkiLCJhdF9oYXNoIjoiNkp2bE9JcmwxSWt0b2U4MTY0ODd2USIsImlhdCI6MTY3MjE5Nzc2OSwiZXhwIjoxNjcyMjAxMzY5fQ.PsKqRS-S8Hs4qYWeZMkA-5YrTXgNYgp3opsEZNfFg7GEvjbEp8H_6xoJRKsj9V8Tmoq0uknHcTXno5sCU-Z1_ZKnnw-VwkVPyAopg5ulB10o4IjXSlIPx-K685WcieUvoZHc-_Ya4y4ZGp3Q5smBU2ObavCq6gzhl8M615bWjwxrEh9bJ00QL2kkJiduHoeZ0XPS9EpvH9nRirUSwrl3rm5z-yIkU8qPpgBwY2AZksLzBBjPBNcIXtnjQ1_x3H_e4ODpEDVQQYxCrYD74WVDY7XwZcWqOTs3Ar7FDpf8tb3csxDsNnew_C2qGkned7FpnvwfCZwRmj0L5a9Jl-QZMg"

	dataArray := strings.Split(idToken, ".")
	// IDトークンからヘッダー・ペイロード・署名を分割
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64デコードする(jsonに戻す)
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64デコードする(jsonに戻す)
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 公開鍵構造体を作る
	E := "AQAB"
	N := "nnnnnnnn"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// 検証するデータ
	// ヘッダー＋ペイロードをハッシュ化して、メッセージダイジェストを生成
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名をbase64デコード（jsonに変換する）
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// 比較検証する
	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
