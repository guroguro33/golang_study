package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// config.iniファイルを生成し、変数を定義する
// 変数の構造体を作成
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// 定義した構造体の変数を宣言
var Config ConfigList

func init() {
	// configファイルを読み込み
	cfg, _ := ini.Load("config.ini")
	// 値を読み込み
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"), // Must系の場合、第１引数に初期値を設定
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
