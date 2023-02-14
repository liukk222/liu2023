package main

import "github.com/gin-gonic/gin"

func TestJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "liukk222",
		"site": "https://github.com/liukk222",
	})
}

func TestXML(c *gin.Context) {
	c.XML(200, gin.H{
		"name": "liukk222",
		"site": "https://github.com/liukk222",
	})
}

func TestHtml(c *gin.Context) {
	c.HTML(200, "test1.5.html", nil)
}

func TestString(c *gin.Context) {
	c.String(200, "hello,wrod")
}
func main() {
	e := gin.Default()
	e.GET("/test_json", TestJson)
	e.GET("/test_xml", TestXML)
	e.LoadHTMLGlob("templates/*")
	e.GET("/test_html", TestHtml)
	e.GET("/test_string", TestString)
	e.Run(":8888")
}
