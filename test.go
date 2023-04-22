package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	username := "your_username"
	password := "your_password"
	url := "https://api.example.com"

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 添加Basic认证头
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.Status)
}
