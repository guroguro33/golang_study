package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey string, sign string, data []byte) {
	apiSecret := DB[apiKey]
	// 秘密鍵を使ってハッシュ化させる
	h := hmac.New(sha256.New, []byte(apiSecret))
	// ハッシュにdataを追加
	h.Write(data)
	// nilを追加してhexにエンコードする
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	// エンコードしたexpectedHMACとクライアント側でハッシュ化したsignを比較
	fmt.Println(sign == expectedHMAC) // true
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	// ハッシュ化したいメッセージdataをbyteスライスにして準備
	data := []byte("data")
	// hmac.New(sha256.New, key []byte)でhmacによるハッシュ化を行う
	h := hmac.New(sha256.New, []byte(apiSecret))
	// dataをハッシュに追加する
	h.Write(data)
	// nilを追加してhexにエンコードする
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign) // 出力:80f11c0d1c2c4c7205d11a9be6bff199371b423ad07cffa7826fff1273f7e449

	// サーバーとハッシュが一致するかチェック
	Server(apiKey, sign, data)
}
