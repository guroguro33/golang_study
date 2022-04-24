package mylib

import "fmt"

var Public string = "Public"   // 大文字変数のため、public
var private string = "private" // 小文字変数のためprivate

// typeやclass名が大文字だとpublic、小文字だとprivateとなる
type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
