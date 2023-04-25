package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	dom, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text := s.Text()
		fmt.Println(i, href, text)
	})

}
