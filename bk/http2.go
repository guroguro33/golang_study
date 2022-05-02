package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 正しいURLかparse（解析）する
	base, _ := url.Parse("http://example.com")
	// URLの後半部分を生成
	reference, _ := url.Parse("/test?a=1&b=2")
	// アクセスするエンドポイント生成(ResolveReferenceを使用)
	endpoint := base.ResolveReference(reference).String()
	// fmt.Println(endpoint)

	// リクエストを生成
	req, _ := http.NewRequest("GET", endpoint, nil) // getなので第３引数はnil
	req.Header.Add("If-None-Match", `test`)
	q := req.URL.Query()
	fmt.Println(q)

	// httpクライアントを作成し、実際にアクセスし、レスポンスを確認する
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
