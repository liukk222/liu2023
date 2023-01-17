package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func MyServer() {
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world!")
	}
	http.HandleFunc("/hello", f)
	http.ListenAndServe(":9999", nil)
}

type countHandler struct {
	mu sync.Mutex
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "conut is %d", h.n)
}
func testHttpServer2() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	testHttpServer2()
	//MyServer()
	/* var test *string
	fmt.Println(test)
	test = new(string)
	*test = "测试"
	fmt.Println(*test) */
}
