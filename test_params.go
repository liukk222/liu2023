package main

import "github.com/gin-gonic/gin"

func TestGet(c *gin.Context) {
	s := c.Query("username")
	pwd := c.DefaultQuery("password", "123")
	c.String(200, "username:%s,password:%s", s, pwd)
}
func TestPathParam(c *gin.Context) {
	s := c.Param("name")
	s2 := c.Param("age")
	c.String(200, "name:%s, age:%s", s, s2)
}
func GoSerarch(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}
func Serach(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	key := c.PostForm("key")
	c.String(200, "page:%s,key:%s", page, key)
}
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	//e.GET("/testGet", TestGet)
	//e.GET("/testpath/:name/:age", TestPathParam)
	e.GET("/goserarch", GoSerarch)
	e.POST("/Serach", Serach)
	e.Run(":8888")
}
