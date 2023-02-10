package main

import "github.com/gin-gonic/gin"

func Go(c *gin.Context) {
	c.HTML(200, "test_static.html", nil)
}

func main() {
	e := gin.Default()
	e.Static("/assets", "./assets")
	e.LoadHTMLGlob("templates/*")
	e.GET("/go", Go)
	e.Run(":8888")
}
