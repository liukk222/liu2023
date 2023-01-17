package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func test1() {
	name := "liukk222"
	templateStr := "Hello,{{.}}"
	t := template.New("test")
	t2, err := t.Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}
	t2.Execute(os.Stdout, name)
}
func test2() {
	type Person struct {
		Name string
		Age  int
	}
	liu := Person{Name: "liukk222",
		Age: 19}
	muban := "hello,{{.Name}},your age is{{.Age}}"
	tmpl, err := template.New("test2").Parse(muban)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(os.Stdout, liu)
}
func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("test.html")
	if err != nil {
		log.Fatal(err)
	}
	t1.Execute(w, "hello,world")
}

func testHTML() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}
func main() {
	testHTML()

}
