package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type LoginUser1 struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type Customer struct {
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
	t, err := engine.DBMetas()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range t {
		// fmt.Printf("v: %v\n", v.Name)
		t2, err2 := engine.TableInfo(v)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Printf("t2: %v\n", t2)

	}
}

func main() {
	engine.Sync(new(LoginUser1), new(Customer))
}
