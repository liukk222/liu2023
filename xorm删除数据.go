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
func test1() {
	engine.ShowSQL(true)
	user := new(User2)
	affected, _ := engine.ID(1).Delete(user)
	fmt.Printf("affected: %v\n", affected)
}

func main() {
	user := User2{}
	// engine.Sync2(&user)
	// user.Name = "liu"
	// user.Age = 19
	// user.Passwd = "123"
	//engine.Insert(&user)
	engine.ID(1).Delete(&user)            //软删除
	engine.ID(1).Unscoped().Delete(&user) //物理删除
}
