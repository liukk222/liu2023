package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.String(200, "hello,%s", "liu")
}
func MyMiddleware1(c *gin.Context) {
	fmt.Println("我的第一个中间件")
}
func MyMiddleware2(c *gin.Context) {
	fmt.Println("我的第二个中间件")
}
func main() {
	e := gin.New()
	e.Use(MyMiddleware1, MyMiddleware2)
	e.GET("/go", test)
	e.Run(":8888")
}
