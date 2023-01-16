package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testGet() {

	url := "https://www.baidu.com"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Println("err2")
		log.Fatal(err2)
	}
	fmt.Printf("b: %v\n", string(b))
}

func main() {
	testGet()
}
