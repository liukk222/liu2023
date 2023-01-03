package main

import (
	"fmt"
	"log"
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

func MyTransactionOps() error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		fmt.Println("err1")
		log.Fatal(err)
	}

	user1 := User2{Name: "liu", Age: 19, Passwd: "123"}
	if _, err := session.Insert(&user1); err != nil {
		fmt.Println("err2")
		log.Fatal(err)
	}
	user2 := User2{Name: "liu", Age: 19, Passwd: "123"}
	if _, err := session.Insert(&user2); err != nil {
		fmt.Println("err2")
		log.Fatal(err)
	}
	//session.Rollback()

	// add Commit() after all actions
	return session.Commit()

}

func main() {
	MyTransactionOps()
}
