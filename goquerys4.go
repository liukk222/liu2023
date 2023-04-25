package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		fmt.Printf("href: %v\n", href)
		detail_url := url + href
		d, _ := goquery.NewDocument(detail_url)
		s2 := d.Find(".article-title").Text()
		fmt.Printf("s2: %v\n", s2)
		ret, err := d.Find(".article").Html()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("ret: %v\n", ret)
	})
}
