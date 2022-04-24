# Go学習メモ

# 環境

## コマンド

```go
// 実行
go run main.go
// フォーマットチェック
gofmt main.go
// フォーマット実行
gofmt -w main.go
```

# Go doc（ドキュメント）

```go
// Printfの解説を見る
go doc fmt.Printf
```

# Printfの使い方

```go
fmt.Printf("%v\n", val)  // value
fmt.Printf("%#v\n", val) // syntax付き値
fmt.Printf("%T\n", val)  // type
fmt.Printf("%t", val)    // boolean
fmt.Printf("%d", val)    // 10進数
fmt.Printf("%s", val)    // 文字列
```

# 文法

## フォーマット

```go
package main

import (
	"fmt"
)

// 初期化関数init（mainより先に呼ばれる）
func init() {
	// ソース　
}

func main() {
	fmt.Println("Hello, world!!", "カンマで連結できるよ")
}
```

## 変数

```go
// **関数内**のみ有効な変数（型は自動的に決まる）
	city := "tokyo"
	tax  := 0.1 // 型はfloat64になる 
```

```go
// 複数変数の宣言
	var nickname, country string
// 複数変数の代入
	nickname, country = "JIRO", "JAPAN"
```

```go
// 複数変数の宣言その２
	var (
		test1 string = "複数変数１" // 型指定も可能
		test2        = "複数変数２"
	)
```

```go
// 宣言のみの初期値
var (
	i int       // 0が初期値
	f64 float64 // 0が初期値
	s string    // ""が初期値
	t, f bool   // falseが初期値
)
```

# 定数

```go
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)
```

## 型

### int型

```go
// x = x + 1
x++
// ++xはできない

// ビット演算子（PHPと一緒）
1 << 0 // 0001 が 0001
1 << 1 // 0001 が 0010
1 << 2 // 0001 が 0100
```

### string型

```go
fmt.Println("Hello world")
fmt.Println("Hello" + " world")      // 文字列結合は+
fmt.Println("Helloworld"[0])         // ASCIIコードの72が出力
fmt.Println(string("Helloworld"[0])) // ASCIIコードの72をキャストしてHが出力
var s string = "Hello world"
s = strings.Replace(s, "H", "X", 1) // sの中からHをXに置き換える（１回だけ）
fmt.Println(s)
fmt.Println(`TEST
		TEST
TEST`) // 改行などが有効になる
```

### bool型

```go
t, f := true, false
fmt.Printf("%T %v\n", t, t) // bool true
fmt.Printf("%T %v\n", f, f) // bool false

// && || などは多言語と一緒
```

### キャスト

```go
var x int = 1
xx := float64(x)
fmt.Printf("%T %v %f\n", xx, xx, xx) // float64 1 1.000000

var y float64 = 1.2
yy := int(y)
fmt.Printf("%T %v %d\n", yy, yy, yy) // int 1 1

var s string = "14"
i, _ := strconv.Atoi(s)     // そのままキャストできないので、パッケージstrconvを使う
fmt.Printf("%T %v\n", i, i) // int 14
```

### バイト型

```go
b := []byte{72, 73}
fmt.Println(b)         // [72 73]
fmt.Println(string(b)) // HI

c := []byte("HI")
fmt.Println(c)         // [72 73]
fmt.Println(string(c)) // HI
```

# 配列

サイズを指定するので、サイズ変更できない

```go
// 変数宣言と初期化
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a) // [100 200]
```

# スライス

サイズを指定しないので、サイズ変更できる

```go
// 変数宣言と初期化
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b) // [100 200 300]

// スライス
	n1 := [...]int{1, 2, 3, 4, 5}
	n2 := n1[1:3] // 1番目から2番目までを切り取る（参照）
	n2[0] = 99 // 切り取った1つ目の配列を99にする（参照変数なので元から変わる）

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])   // 3
	fmt.Println(n[2:4]) // [3 4] カンマの位置で2個目から４個目まで
	fmt.Println(n[:2])  // [1 2] カンマの位置で0個目から４個目まで
	fmt.Println(n[2:])  // [3 4 5 6] カンマの位置で2個目以降
	fmt.Println(n[:])   // [1 2 3 4 5 6]

// 配列の中に配列を入れる
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board) // [[1 2 3] [4 5 6] [7 8 9]]
```

### make関数

第２引数のlengthに値を入れると、0埋めされる

```go
	n := make([]int, 3, 5)                                            // [0 0 0]
	fmt.Printf("length=%d capacity=%d value=%v\n", len(n), cap(n), n) // length=3 capacity=5 value=[0 0 0]

	n = append(n, 0, 0)                                               // [0 0 0 0 0]
	fmt.Printf("length=%d capacity=%d value=%v\n", len(n), cap(n), n) // length=5 capacity=5 value=[0 0 0 0 0]

	a := make([]int, 3)                                               // lengthだけ指定し、capacityは指定しない
	fmt.Printf("length=%d capacity=%d value=%v\n", len(a), cap(a), a) // length=3 capacity=3 value=[0 0 0]

	b := make([]int, 0)
	var c []int
	fmt.Printf("length=%d capacity=%d value=%v\n", len(b), cap(b), b) // length=0 capacity=0 value=[] nilを入れてメモリ使用
	fmt.Printf("length=%d capacity=%d value=%v\n", len(c), cap(c), c) // length=0 capacity=0 value=[] メモリ不使用
```

# マップ（連想配列）

```go
// map[キーの型]値の型　で定義する
var m map[string]int{"a": 1, "b": 2}

var m = map[string]int{"apple": 100, "banana": 200}
fmt.Println(m)
fmt.Println(m["banana"]) // 200

// 要素の書き換え
m["banana"] = 300
fmt.Println(m) // map[apple:100 banana:300]

// 要素の追加
m["new"] = 500
fmt.Println(m) // map[apple:100 banana:300 new:500]

// ない要素を指定すると0になる
fmt.Println(m["nothing"]) // 0

// 戻り値の２つ目に存在判定がtrueかfalseで返る
v, ok := m["apple"]
fmt.Println(v, ok) // 100 true

v2, ok2 := m["nothing"]
fmt.Println(v2, ok2) // 0 false
```

# 関数

```go
func add(x int, y int) int { // 引数と返り値に型指定
	return x + y
}

func add(x, y int) int {     // 複数の引数の型をまとめて定義
	return x + y
}

func add(x, y int) (int, int) {  // 返り値が複数の場合は型にparenthesisの括弧が必要
	return x + y, x * y
}

func cal(price, item int) (result int) { // 返り値に変数を定義する
	result = price * item                  // resultは定義済みのため:=ではない
	return result                          // このresultは省略可能
}

// 関数を変数に代入して使うことも可能
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(100)

// 即時関数もOK
	func(x int) {
		fmt.Println("inner func", x)
	}(200)
```

## クロージャ

- 関数を返却することによって、それ以前の計算は１度だけ実行される
- 「グローバル変数の節約」と「無駄な計算を何度も行わない」が実現できる

```go
func incrementGenerator() (func() int) {
	x := 0  // この計算は呼び出し時の１度しか実行されない、ここで計算を１度だけさせ、変数の内容を維持できる
	return func() int {
		x++
		return x
	}
}

func main() {
	counter := incrementGenerator() // クロージャ関数のreturn前の部分を１度だけ計算する
	fmt.Println(counter())          // 1
	fmt.Println(counter())          // 2
	fmt.Println(counter())          // 3
}
```

## 可変長引数

- ...を使用する

```go
func foo(params ...int) {
	fmt.Println(len(params), params) // 長さと値を出力
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...) // スライスを展開して引数に使用する
}
```

# if文

- 基本は他言語と一緒
- 違いは条件文の括弧が不要

```go
// ローカル変数を作りつつ、判定を行うもの
func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "NO"
	}
}

func main() {
	if result := by2(10); result == "OK" { // resultという変数を作成しつつ、判定に使う
		fmt.Println("great")
	}
	// fmt.Println(result) ローカル変数のため、外から呼び出し不可
}
```

# for文

- 基本は他言語と一緒(foreachやwhileはgoに構文なし)
- 違いは、初期値設定とインクリメントは省略が可能
- continueやbreakは使用可能

```go
// 通常
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

// 初期化とインクリメントを省略（セミコロンも省略可能）
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
```

## for range文

- keyとvalueを取り出す際に使える便利なrange

```go
	slice := []string{"python", "go", "java"}

	for key, val := range slice {
		fmt.Println(key, val)
	}
	// 0 python
	// 1 go
	// 2 java

// keyが不要の時は_で握りつぶす
	for _, val := range slice {
		fmt.Println(val)
	}
	// python
	// go
	// java

// mapの場合
	m := map[string]int{"apple": 100, "banana": 200}

	for key, val := range m {
		fmt.Println(key, val)
	}
	// apple 100
	// banana 200

// keyだけでいいときはsliceと違い、_不要
	for key := range m {
		fmt.Println(key)
	}
	// apple
	// banana

// valだけ
	for _, val := range m {
		fmt.Println(val)
	}
	// 100
	// 200
```

# switch文

- 条件変数部分に括弧が不要、breakが不要
- 条件変数部分をなしでもOK、case部分に条件を記載する

```go
// 通常パターン
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}

// 条件部分にローカル変数を使用
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}

// 条件部分なしでcaseに条件式を書く
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}

	func getOsName() string {
		return "mac"
	}
```

# defer

- 遅延実行させる

```go
// deferをつけたものを最後に実行する
	defer fmt.Println("world")
	fmt.Println("Hello")
	// Hello
	// world

// LIFO（スタックしたものを最初に実行）
	fmt.Println("run")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("success")
	// run
	// success
	// 3
	// 2
	// 1

// 使用例
	file, _ := os.Open("./main.go")
	defer file.Close() // Closeを忘れないようにする
	// byte配列を100バイト分用意する
	data := make([]byte, 100)
	// byte配列にfileの中身を入れる
	file.Read(data)
	fmt.Println(string(data))
```

# log

- よくあるinfoやwarningなどのpakageは標準ライブラリはなし。使いたい場合はサードパーティのものを利用する。

```go
// ログ出力
	log.Println("logging!!")                // 2022/03/14 05:12:46 logging!!

// Fatal系のメソッドは実行するとexitする
	log.Fatalf("%T %v", "Fatal!", "Fatal!") // 2022/03/14 05:12:46 string Fatal!
	log.Fatalln("logging!!Fatal!")          // exitされて実行されない

// 使用例
	_, err := os.Open("./fwefwffwf")
	if err != nil {
		log.Fatalln("Exit!", err) // 2022/03/14 05:16:40 Exit! open ./fwefwffwf: no such file or directory
	}

// main関数内でログ設定を実行
	LoggingSettings("test.log")

func LoggingSettings(logFile string) {
	// ファイルを読み込む際に読み書きOKにする、ファイルが存在しなかったら新規作成する
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// multiwriterで画面とログファイルの両方に出力する
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// Llongfileでフルパス名が出力される
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// SetOutputで出力先を変更
	log.SetOutput(multiLogFile)
}
```

# エラーハンドリング

try-catchがないので、下記の様にエラー処理を行う

```go
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data) // errを２回宣言しているが、複数宣言のどちらかが成功すれば宣言可能（errは上書きされる）
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))
```

# panicとrecover

- panic関数で処理を停止
- recover関数で処理を復帰

```go
func thirdPartyConnectDB() {
	panic("Unable to connect database!") // panicで処理停止する
}

func save() {
	defer func() {
		s := recover() // recoverでpanicしても処理を停止しない
		fmt.Println(s) // panicの中身を出力
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
```

# ポインタ

アドレスとメモリ

```go
	var n int = 100
	fmt.Println(n)  // 100

	fmt.Println(&n) // &アンパサンドをつけるとアドレスになる 0x0004

// *int ポインタ型
	var p *int = &n // アドレスを別のアドレスのメモリに定義する 
	fmt.Println(p)  // メモリに格納したアドレスを表示 0x0004
	fmt.Println(&p) // 別のアドレスのアドレスを表示   0x0004
	fmt.Println(*p) // アドレスの値を表示           100
```

![スクリーンショット 2022-03-15 6.10.33.png](Go%E5%AD%A6%E7%BF%92%E3%83%A1%E3%83%A2%206aa28235f53f499fa642cc6eefc37db1/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88_2022-03-15_6.10.33.png)

# newとmakeの違い

ポインタ型を返すものはnew、そうでないものはmake

```go
// newで初期化
	var p *int = new(int)
	fmt.Println(*p) // 0
	*p++
	fmt.Println(*p) // 1

// 宣言のみ
	var p2 *int
	fmt.Println(p2) // nil
	*p2++           // nilに対してインクリメントできない

// 配列やmapはmake
	s := make([]int, 0)
	fmt.Printf("%T %v\n", s, s) // []int []

	m := make(map[string]int)
	fmt.Printf("%T %v\n", m, m) // map[string]int map[]

	ch := make(chan int)          // チャンネル
	fmt.Printf("%T %v\n", ch, ch) // chan int 0xc0000240c0

// ポインタやstructはnew
	var p *int = new(int)       // ポインタ
	fmt.Printf("%T %v\n", p, p) // *int 0xc0000140a0

	var st = new(struct{})
	fmt.Printf("%T %v\n", st, st) // *struct {} &{}
```

# struct（構造体）

フィールドの集まり

```go
// 基本
// structの定義
type Vertex struct {
	X int // 大文字(capital)
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2, S: "test"}
	fmt.Println(v) // {1 2 test}

	// .をつけて呼び出す
	fmt.Println(v.X, v.Y, v.S) // 1 2 test
	v.X = 100
	fmt.Println(v) // {100 2 test}

	// 未指定は初期化
	v2 := Vertex{X: 20}
	fmt.Println(v2) // {20 0 (空文字)} intは0、stringは空文字で初期化される

	v3 := Vertex{}
	fmt.Printf("%T %v\n", v3, v3) // main.Vertex {0 0 空文字}

	// structを以下の形で宣言時、他の型と違い、nilではないので注意
	var v4 Vertex
	fmt.Printf("%T %v\n", v4, v4) // main.Vertex {0 0 空文字}

	v5 := new(Vertex)
	fmt.Printf("%T %v\n", v5, v5) // *main.Vertex &{0 0 } ポインタ型でアドレスが変える

	// 上よりこっちの方がポインタとわかりやすい
	v6 := &Vertex{}
	fmt.Printf("%T %v\n", v6, v6) // *main.Vertex &{0 0 } ポインタ型でアドレスが変える
}

// 応用
func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 2000
	// (*v).X = 2000 本来はポインタで実体を指定するが、goはよしなにしてくれる
}

func main() {
	v7 := Vertex{1, 2, "test"}
	changeVertex(v7) // 別ポインタで数値を変えるので、v7は変化しない
	fmt.Println(v7)

	v8 := &Vertex{1, 2, "test"}
	changeVertex2(v8) // アドレスで渡すので、v8が書き変わる
	fmt.Println(v8)   // &{2000 2 test}
	fmt.Println(*v8)  // {2000 2 test}
}
```

# オブジェクト指向（Object-Oriented）のようなもの

## ポインタレシーバーと値レシーバー

- GOにクラスが、レシーバーが似た様なことをする

```go
// 構造体定義
type Vertex struct {
	X, Y int
}

// 値レシーバー（値vを渡している）これでvがAreaメソッドを使えるよう紐付けをした
func (v Vertex) Area() int {
	return v.X * v.Y
}

// ポインタレシーバー（ポインタを渡している）vがScaleメソッドを所有
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	// 構造体
	v := Vertex{6, 8} // インスタンス作成のようなもの
	v.Scale(10) // ポインタレシーバーにより、vが10倍された
	fmt.Println(v.Area())

}
```

## コンストラクタ

構造体定義の変数を大文字から小文字に変更すると、外部から参照できなくなる

その際に、Newメソッドを作成し、外部から使えるコンストラクタとする（setter的なもの？）

```go
type Vertex struct {
	// 変数を小文字にするとpackage内からの変更が可能になり、package外から変更できない
	x, y int
}

// コンストラクタ
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
}
```

# Embedded（埋め込み）

GOに継承はないが似たことができる

type宣言時に、上位typeを呼び出す

```go
// 上記のコンストラクタソースがあるとして、、、
type Vertex3D struct {
	Vertex // Vertexのxとyを埋め込んで新しいtypeを作成
	z int
}

// 値レシーバー（値vを渡している）
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// ポインタレシーバー（ポインタを渡している）
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// コンストラクタ
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z} // 構造体の中に構造体を記述
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D()) // 60000

}
```

# non-struct（非構造体）

```go
// int型の独自intを定義（非構造体）
type Myint int

// 非構造体でもメソッドを登録できる
func (i Myint) Double() int {
	fmt.Printf("%T %v\n", i, i) // main.Myint 10
	return int(i * 2) // Myintを使った計算なのでintでcastする
}

func main() {
	myInt := Myint(10)
	fmt.Println(myInt.Double()) // 20
}
```

# インターフェース

```go
// インターフェースを定義
type Human interface {
	// Sayメソッドが必須となる
	Say() string
}

type Person struct {
	Name string
}

// PersonクラスにSayメソッドを定義
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func main() {
	// Human型mikeはSayメソッドがあるのでOK
	var mike Human = &Person{"Mike"}
	mike.Say() // Mr.Mike
}
```

# ダッグタイピング

```go
// インターフェースのダッグタイピング(Sayが必須なHuman)
func DriveCar(human Human) {
	// humanはSayメソッドが必須なのでOK
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}
```

# タイプアサーション

```go
// タイプアサーション
// キャストはGoだとタイプコンバージョンと言う
func do(i interface{}) {
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

func main() {
	do(10) // 20
}

// スイッチタイプ文を使い、型で分岐を行う
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)     // 20
	do("test") // test!
	do(true)   // I don't know bool
}
```

# Stringer

Stringerはfmtパッケージ内に用意されたインターフェースであり、Stringメソッドを持っている

interfaceに定義された関数をメソッドとして定義すると暗黙的にinterfaceを実装したことになる

```go
// fmtパッケージ内の記述
type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

// String()メソッドを定義して出力内容を変更すると、こちらを優先して呼び出しを行う
func (p Person) String() string {
	return fmt.Sprintf("My name is %v\n", p.Name)
}

func main() {
	mike := Person{"Mike", 23}
	fmt.Println(mike) // My name is Mike
}
```

# カスタムエラー

独自のエラー構造体を作成し、Errorメソッドの出力をカスタマイズする

```go
// 独自のエラーを作成
type UserNotFound struct {
	Username string
}

// Errorメソッドをオーバーライドして、独自Error出力を作成
func (e *UserNotFound) Error() string { // 独自エラーをアドレス受けする
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"} // 参照渡しがお約束
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
```

# goroutine

```go
func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// goをつけて実行すると並列実行される
	go goroutine("world")
	normal("hello")
}

// 実行結果
hello
world
world
hello
hello
world
world
hello
world
hello
```

# sync.WaitGroup

上記の例ではtime.Sleepで待機時間がないと、normal関数の実行が終わり次第、goroutine関数は実行される前に完了してしまう

そのため、sync.WaitGroupを設定して、並列処理が完了するまで待たせることができる

```go
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	// wg.done()を実行して終了を知らせる
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	// sync.WaitGroupを定義し、１つの関数があるためAdd(1)する
	var wg sync.WaitGroup
	wg.Add(1)
	// goをつけて実行すると並列実行される
	go goroutine("world", &wg)
	normal("hello")
	// wg.done()が１回あるまで処理を待つ
	wg.Wait()
}
```

# Channel

goroutineとのデータの受け渡しを行うためにchannelを使う

```go
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 結果をchannelに渡す
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 結果をchannelに渡す
	c <- sum
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	// channelを定義して並列処理の関数に渡す
	c := make(chan int)
	go goroutine1(s1, c)
	go goroutine2(s2, c)
	// チャンネルから値を取り出す(channelにはキューで出し入れするため、早く入った方から出力される)また、ブロッキングで待機する
	x := <-c
	fmt.Println(x) // 40
	y := <-c
	fmt.Println(y) // 15
}
```

# Buffered Channels

channelをmakeするときは、第２引数にバッファの長さを指定する

指定数以上にchannelの中に値を入れられない

そのままfor-range関数で回すとエラーになるため、close(ch)でクローズしてあげる必要あり

```go
	// バッファの長さを2とする
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch)) // 1
	ch <- 200
	fmt.Println(len(ch)) // 2
	// closeしてやらないと、range関数で3つ目以降のchannelを取り出そうとしてエラーになる
	close(ch)

	for c := range ch {
		fmt.Println(c) // 100 200
	}
```

# Channelのrangeとclose

for-rangeでchannelを回すときはcloseが必要

```go
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	// forが終わったらcloseしないと、for-rangeで次を取りに行ってエラーになる
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s)) // バッファを指定
	go goroutine1(s, c)
	// forはgoroutine1の実行待ちをして、channelに値が入ったら実行する
	for i := range c {
		fmt.Println(i)
	}
}
```

# producerとcomsumer

イメージは下記の通り

- goroutineのproducerで値を取得してくる
- goroutineのconsumerで値を処理する

```go
func producer(ch chan int, i int) {
	// producerではapiから取得などの処理を記述
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("proccess", i*1000)
		wg.Done() // 10回分のAddをここでDoneさせている
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg) // 内部でproducerがchに入れたものを都度出力して、wg.Done()している
	wg.Wait()　　　　　　　// 全てDoneするまで待機　　　　　　　　　
	// main関数ではないgoroutineのconsumer関数内でchannelのrangeループをしており、chに値が入ってくるのをずっと待機しているため、closeしてやる必要あり
	close(ch)
}
```

# fan-out fan-in

![スクリーンショット 2022-03-29 5.30.05.png](Go%E5%AD%A6%E7%BF%92%E3%83%A1%E3%83%A2%206aa28235f53f499fa642cc6eefc37db1/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88_2022-03-29_5.30.05.png)

非同期でいくつもの関数にchanを渡しながら処理を続けていく

```go
func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// channel定義の際に、<-chanとchan<-のように送信と受信を表示するとわかりやすい（なくても動く）
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)
	// 待機しながらchanに値が入ると出力している
	for result := range third {
		fmt.Println(result)
	}
}
```

# channelとselect-case

select文を使うことで、ゴルーチンは複数の操作を待機する

- selectは、ケースの1つが実行可能になるまで他のケースをブロックし、その後に他ケースを実行します。
- 複数のケースが準備ができている場合はランダムに1つを選択して順に処理して行きます。

```go
func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine1(c1)
	go goroutine2(c2)
  // select内で待機する 
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
```

# Default Selectionとfor break

- select case文では該当がなければdefaultに移行する
- forの中にselect文を書く際は、returnでmain関数から抜けたり、break OuterLoopなどでfor文から抜ける

```go
func main() {
	tick := time.Tick(100 * time.Millisecond)  // chanを返す
	boom := time.After(500 * time.Millisecond) // chanを返す

OuterLoop:
	for {
		select {
		case <-tick: // chanからの代入がなくてもOK
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop // OuterLoopが終了（）breakだけだとselect文だけ抜ける
			// return // returnでmain関数終了もできる
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

# sync.Mutex

排他制御することができるので、書き込み競合などのエラーを防ぐことができる

sync.Mutexで型定義し、Lock()とUnLock()で排他制御を行う

```go
type Counter struct {
	v   map[string]int
	mux sync.Mutex // Mutexを定義する
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()         // 排他ロック
	defer c.mux.Unlock() // 排他ロック解除
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main()
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("key")) // 20
	fmt.Println(c)              // {map[key:20] {0 0}}
}
```

# パッケージ

フォルダ名がパッケージ名になる

パッケージ名＋メソッド名で呼び出し

```go
import (
	// abc順はgofmt
	// 標準ライブラリ->自作~等の並び順になるのがgoimports
	"fmt"

	"awesomeProject/mylib"       // フォルダ名がパッケージ名になる
	"awesomeProject/mylib/under" // 階層も持てる
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s)) // ファイル名 != メソッド名なので間違えずにメソッド名を呼ぶ
	mylib.Say()
	under.Hello()
}
```

# PublicとPrivate

変数名、型名、メソッド名の最初を大文字にするとPublicになり、小文字にするとPrivateになる

```go
// package
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

// main.go
package main

import (
	"awesomeProject/mylib"       // フォルダ名がパッケージ名になる
)

func main() {
	person := mylib.Person{Name: "taro", Age: 25}
	fmt.Println(person)

	fmt.Println(mylib.Public)
}

```

# Test

テストはパッケージファイルと同じ階層にxxx_test.goファイルを以下の様に作成する

```go
package mylib

import "testing" // testingを必ずインポートする

// TestXxxx(t *testing.T)という関数名にする
func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}

// 実行コマンド
go test -v ./... // -vオプションで結果の詳細を表示　./...で今の階層以下の全てのテストを探して実行
```

# 標準パッケージ

## time

```go
func main() {
	// 現在時刻取得
	t := time.Now()
	fmt.Println(t) // 2022-04-18 05:49:08.931603 +0900 JST m=+0.000135271

	// PostgreSQLで代入可能なRFC3339の形にする
	f := time.Now().Format(time.RFC3339)
	fmt.Println(f) // 2022-04-18T05:49:08+09:00

	// 時間など抽出する場合
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}
```

## regex

```go
func main() {
	// 1行で正規表現とチェックしたい値を記述（正規表現を１回しか使わない）
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match) // true

	// 正規表現を何度も使う...MustCompileしてからMatchString
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms) // true

	// 一致すれば文字列を返す...FindString
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs) // /edit/test

	// 一致すれば文字列と/区切りの文字列を配列で返す...FindStringSubmatch
	fss := r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2]) // [/edit/test edit test] /edit/test edit test
}
```

## sort

```go
func main() {
	i := []int{5, 3, 2, 8, 6}
	s := []string{"d", "a", "k"}
	// 1度しか使わないstructは直接定義できる
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}

	fmt.Println(i, s, p) // 結果　[5 3 2 8 6] [d a k] [{Nancy 20} {Vera 40} {Mike 30} {Bob 50}]

	// intの並び替え
	sort.Ints(i)
	// stringの並び替え
	sort.Strings(s)
	// structのNameでの並び替え
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	// structのageでの並び替え
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })

	fmt.Println(i, s, p) //　結果　[2 3 5 6 8] [a d k] [{Nancy 20} {Mike 30} {Vera 40} {Bob 50}]
```