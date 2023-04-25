package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("请求前调用:OnRequest")
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("发生错误调用:OnError")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("获得响应后调用:OnResponse")
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("OnResponse收到html内容后调用:OnHTML")
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println("OnResponse收到xml内容后调用:OnXML")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("结束", r.Request.URL)
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
}
