package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 编码字符串
	encoded := url.QueryEscape("https://www.example.com/search?q=golang&sort=recent")
	fmt.Println("Encoded:", encoded)

	// 解码字符串
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decoded:", decoded)
}
