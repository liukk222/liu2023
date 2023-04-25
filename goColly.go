package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url", r.URL)
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
}
