package main

import (
	// abc順はgofmt
	// 標準ライブラリ->自作~等の並び順になるのがgoimports
	"fmt"

	"awesomeProject/mylib"       // フォルダ名がパッケージ名になる
	"awesomeProject/mylib/under" // 階層も持てる
)

func main_bk() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s)) // ファイル名 != メソッド名
	mylib.Say()
	under.Hello()

	person := mylib.Person{Name: "taro", Age: 25}
	fmt.Println(person)

	fmt.Println(mylib.Public)
}
