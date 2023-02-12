package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	UId  int    `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make([]User, 3)

func init() {
	u1 := User{1, "tom", 20}
	u2 := User{2, "kite", 30}
	u3 := User{3, "rose", 40}
	users = append(users, u1)
	users = append(users, u2)
	users = append(users, u3)
	fmt.Println(users)
}
func find(uid int) (*User, int) {
	for i, u := range users {
		if u.UId == uid {
			return &u, i
		}

	}
	return nil, -1
}
func (d *User) AddUser(c *gin.Context) {
	users = append(users, *d)
	c.JSON(200, users)
}
func Usersl(us User) *User {
	return &us
}

func DelUser(c *gin.Context) {
	uid := c.Param("uid")
	id, _ := strconv.Atoi(uid)
	_, i := find(id)
	users = append(users[:i], users[i+1:]...)
	c.JSON(200, users)
}
func UpdateUser(c *gin.Context) {
	uid := c.Param("uid")
	id, _ := strconv.Atoi(uid)
	u, _ := find(id)
	name := c.Param("name")

	na := strconv.Quote(name)  //nc := strconv.QuoteToASCII(name) 与na := strconv.Quote(name)一样，返回string类型
	var nc string = na         //na=""字符""
	nc = strings.Trim(na, `"`) //去除"",变为"字符"
	u.Name = nc
	//u.Name = "修改的Name"
	c.JSON(200, u) //json中才为"字符"
}
func FindUser(c *gin.Context) {
	uid := c.Param("uid")
	id, _ := strconv.Atoi(uid)
	u, _ := find(id)
	c.JSON(200, u)
}
func main() {
	e := gin.Default()
	a := User{4, "liu", 20}
	as := Usersl(a)
	e.POST("/user/", as.AddUser)          //通过接受者将main中的User{}传入AddUser函数中，方便进行相关操作。
	e.PUT("/user/:uid/:name", UpdateUser) //能在PUT中修改name
	e.GET("/user/:uid", FindUser)
	e.DELETE("/user/:uid", DelUser)
	e.Run(":8888")

}
