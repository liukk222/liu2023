package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User3 struct {
	Username string `uri:"username"`
	Password string `uri:"password"`
}

func TestUriBind(c *gin.Context) {
	var user User3
	err := c.ShouldBindUri(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

type User2 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User2
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

type User struct {
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "test_from.html", nil)
}
func Regsiter(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	c.String(200, "form data:%s", user)
}
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.GET("/register", GoRegister)
	e.POST("/register", Regsiter)
	// http://localhost:8888/testGetBind?username=liu&password=123
	e.GET("/testGetBind", TestGetBind)
	// http://localhost:8080/testGetBind/liu/123
	e.GET("/testGetBind/:username/:password", TestUriBind)
	e.Run(":8888")

}
