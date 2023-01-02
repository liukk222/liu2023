package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:2003925liu@/liu_xorm?charset=utf8")
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
func created() {
	// 创建表
	err3 := engine.Sync(new(User))
	fmt.Printf("err3: %v\n", err3)
}
func instr() {
	user := User{
		Id:     1,
		Name:   "liukk222",
		Salt:   "salt",
		Age:    19,
		Passwd: "123456",
	}
	affected, _ := engine.Insert(&user)
	fmt.Printf("affected: %v\n", affected)
}

func main() {
	instr()
}
