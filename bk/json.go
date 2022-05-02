package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age,string"` // 小文字のageにしつつ、intではなくstringに
	NickNames []string `json:"nickname"`
}

func main() {
	b := []byte(`{"name":"mike", "age":20, "nicknames": ["a", "b", "c"]}`)

	var p Person
	// Unmarshalはjsonを構造体に変換する
	// json.Unmarshal(data []byte, interface{})
	// Unmarshalしてエラーがあったらerrに入る
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.NickNames) // mike 20 [a b c]

	// marshalは構造体をjsonに変換する
	// json.Marshal(interface{})
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
	// 出力：
	// 構造体で指定なし：{"Name":"mike","Age":20,"NickNames":["a","b","c"]}
	// 構造体の最後に`json:"name"など指定：`{"name":"mike","age":"20","nickname":null}

}
