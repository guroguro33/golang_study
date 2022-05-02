package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// http.Getでget通信可能
	resp, _ := http.Get("http://example.com")
	// fmt.Println(resp, err) // errも取得可能
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
