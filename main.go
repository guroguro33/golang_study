package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ioutilでのファイルの操作が可能（ファイル操作に特化） 最新はioとosに移行
	// os.ReadFileがversion1.16以降は推奨されるようだ
	// ファイル操作はosパッケージに移ったようだ
	// content, err := ioutil.ReadFile("main.go") // v1.15まで
	content, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// 簡易文つきif文
	// if 簡易文; 条件 {
	//	条件を満たす場合の式
	//}
	if err := os.WriteFile("os_tmp.go", content, 0666); err != nil { // WriteFile(ファイル名、, 変数(ファイルに書きたいもの), パーミッション)
		log.Fatalln(err)
	}
}
