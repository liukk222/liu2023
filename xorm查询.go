package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type User struct {
	Id      int64
	Name    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
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
func test1() {
	user := new(User)
	id := 1
	has, _ := engine.ID(id).Get(user)
	fmt.Printf("has: %v\n", has)
	fmt.Printf("user: %v\n", user)
}

func test2() {
	user := &User{Id: 1}
	engine.Get(user)
	fmt.Printf("user: %v\n", *user)

}
func test3() {
	engine.ShowSQL(true)
	user := &User{Id: 1}
	b, _ := engine.Exist(user)
	fmt.Printf("b: %v\n", b)
}
func test4() {
	engine.ShowSQL(true)
	user := new(User)
	rows, err := engine.Where("id>?", 1).Rows(user)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.Scan(user)
		fmt.Printf("user.Name: %v\n", user.Name)

	}
}
func main() {
	test4()
}
