package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func save(title string, content string) {
	err := os.WriteFile("./"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	var titles string
	var contents string
	c := colly.NewCollector()

	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
		title := h.Text
		fmt.Printf("title: %v\n", title)
		titles = title
	})

	c.OnHTML(".article", func(h *colly.HTMLElement) {
		content, _ := h.DOM.Html()
		fmt.Printf("content: %v\n", content)
		contents = content
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
	save(titles, contents)
}
