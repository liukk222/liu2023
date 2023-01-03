package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type User2 struct {
	Id        int64
	Name      string
	Age       int
	Passwd    string    `xorm:"varchar(200)"`
	Created   time.Time `xorm:"created"`
	Updated   time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

type Customer struct {
	Id      int64
	Name    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/liu_xorm?charset=utf8")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		err2 := engine.Ping()
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		} else {
			print("连接成功！")
		}
	}
	fmt.Println("init...")
}
func main() {
	user := User2{}
	user.Name = "liu2"
	engine.ID(2).Update(&user)
}
