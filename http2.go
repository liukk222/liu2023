package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://echo.apifox.com/post?q1=%3Cq1%3E&q2=%3Cq2%3E"
	method := "POST"

	payload := strings.NewReader(`<hello>`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (hptts://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("liu", "19")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
