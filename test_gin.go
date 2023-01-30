package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	//c.String(200, "hello,%s", "liukk222")
	c.JSON(200, gin.H{
		"name": "tom",
		"age":  "20",
	})
}
func main() {
	e := gin.Default()
	e.GET("/hello", Hello)
	e.Run()
}
