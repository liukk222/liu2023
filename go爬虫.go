package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Add("Cookie", "__gads=ID=a03841e8dc5e3a01:T=1614315354:S=ALNI_MbPtvxPEZ3zQwJR9ZymNtNXqk5T2w; Hm_lvt_39b794a97f47c65b6b2e4e1741dcba38=1617680418,1617680467,1617691441,1617698379; _ga_R55C9JJH2H=GS1.1.1625526333.1.0.1625526335.0; __utmz=226521935.1627719195.1.1.utmcsr=zzk.cnblogs.com|utmccn=(referral)|utmcmd=referral|utmcct=/; Hm_lvt_1408eee61ba676e12f93290ce7650e83=1629337233; sc_is_visitor_unique=rx11857110.1636164233.13F2C484526C4FA2010957F320FEEA18.1.1.1.1.1.1.1.1.1-9614694.1623747052.1.1.1.1.1.1.1.1.1; __utma=226521935.142504143.1614315356.1627719195.1636780252.2; UM_distinctid=17d4af395843ea-0b45d9f907fcac-978183a-144000-17d4af39585871; _ga_3Q0DVSGN10=GS1.1.1638963943.3.1.1638963943.60; _ga_4CQQXWHK3C=GS1.1.1641964040.3.0.1641964040.0; _ga=GA1.2.142504143.1614315356; Hm_lvt_d8d668bc92ee885787caab7ba4aa77ec=1640569355,1642119762; gr_user_id=3bf443cf-dc53-4d5f-ae5c-5bbb3ca819c3; Hm_lvt_7a41cb9dd1d636656563500edd2ddba8=1646182206; _gid=GA1.2.1450982996.1646736611; .Cnblogs.AspNetCore.Cookies=CfDJ8GsLOKiGtk1Au0UP1SouGdVYsphv8fTJFxTIvJxScUQCqJc5Ugl21LPkwOqhwGAvgS5GW7vDZEpxDA7VMMVyvZdtskQrPLqPj8aNRhFU7bN1vaTnWjRCgmVBKWnkfSOvS71t8xcJFwfWROB6_UEPt9uMWrWdRYlvvInER3kWX2s1rsrDUpUA9HoJ6BaIsnxBv10Xvhixq7gF4187lbmr1ODbLLo8VMRKOUWMrUC3GZHBBRRNP9qLoGvOYLLCwbGfoPEQvbCzXjJTfjM1cLCC0Ajnf4MT3Q-BpwoSmxFKarrunefNYaiVPwGYpJjsxfXFvEQN8rXVlr9MSCcicJepFRs5aQfZZ7z8o2PQomfcn2TZGG8pvdSrCqIESt0fpd9FN3cwwPdqs9aj6MiBEAk4GUeI0_TvTczhW11QHDxyRlFQUtaWaR6JJIcv9xCIC4cMjfOc592R9VjEpdRCqnK0d4NdHsFaDC3UE2SDDjkEmF5qmx7RHdJkPljghmzXC4TdtAX5WhiZMqcV2FJgiH3DjmtPZG0iuSgx9m4qNxYBY7rQpT6JK6MonuNJjOL5LUzbvA; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1646742153,1646796896,1646800170,1646879057; __utmc=59123430; __utmz=59123430.1646882996.44.25.utmcsr=cnblogs.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utma=59123430.142504143.1614315356.1646796901.1646882995.44; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1646883549; __utmb=59123430.3.10.1646882996")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)

}

func pasrse(html string) {
	// 替换掉空格
	html = strings.Replace(html, "\n", "", -1)
	// 边栏内容块正则
	re_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	// 找到边栏内容块
	sidebar := re_sidebar.FindString(html)
	// 链接正则
	re_link := regexp.MustCompile(`href="(.*?)"`)
	// 找到所有链接
	links := re_link.FindAllString(sidebar, -1)

	base_url := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {

		s := v[6 : len(v)-1]
		url := base_url + s
		fmt.Printf("url: %v\n", url)

		body := fetch(url)

		// 启动另外一个线程处理
		go parse2(body)

	}
}

func parse2(body string) {
	// 替换掉空格
	body = strings.Replace(body, "\n", "", -1)
	// 页面内容
	re_content := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	// 找到页面内容
	content := re_content.FindString(body)

	// 标题
	re_title := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	// 找到页面内容
	title := re_title.FindString(content)
	fmt.Printf("title: %v\n", title)
	// 切片
	title = title[42 : len(title)-5]
	fmt.Printf("title: %v\n", title)
	//fmt.Printf("content: %v\n", content)
	save(title, content)
}

func save(title string, content string) {
	err := os.WriteFile("./pages/"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	url := "https://gorm.io/zh_CN/docs/"
	s := fetch(url)
	pasrse(s)

}
