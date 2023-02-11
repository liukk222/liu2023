package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	session := sessions.Default(c)
	//获得session值GET
	if session.Get("hello") != "world" {
		//设置
		session.Set("hello", "liu")
		//保存
		session.Save()
	}

	c.JSON(200, gin.H{"hello2": session.Get("hello")})
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	//注入中间件
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", Handler)

	r.Run(":8888")
}
