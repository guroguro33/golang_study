package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// abcをバイト配列で定義
	r := bytes.NewBuffer([]byte("abc"))
	content, _ := io.ReadAll(r)
	fmt.Println(content)         // [97 98 99]
	fmt.Println(string(content)) // abc
}
