package main

import (
	"fmt"
)

const (
	// const識別子iota（0から始まり、インクリメントしていく）
	c1 = iota
	c2 = iota
	c3 = iota
	c4 // const内で指定がないと直前の呼び出しと同じものが入る
)

const (
	_      = iota
	KB int = 1 << (10 * iota) // 10ビットシフトすると、1024変化する
	MB int = 1 << (10 * iota) // 20ビットシフト
	GB int = 1 << (10 * iota) // 30ビットシフト
)

func main() {
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(KB, MB, GB)
	fmt.Println(1024 * 1024)
}
