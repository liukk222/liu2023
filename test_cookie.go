package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	s, err := c.Cookie("username")
	if err != nil {
		s = "liu"
		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
	}
	c.String(200, "测试cookie")
}

func main() {
	e := gin.Default()
	e.GET("/cookie", Handler)
	e.Run(":8888")

}
