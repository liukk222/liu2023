package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

type LoginUser struct {
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

func main() {

	//engine.ShowSQL(true)
	//engine.SetMapper(names.GonicMapper{})
	//engine.SetMapper(names.SnakeMapper{})
	//engine.SetMapper(names.SameMapper{})
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "golang_")
	engine.SetTableMapper(tbMapper)
	// 创建表
	err3 := engine.Sync(new(LoginUser))
	fmt.Printf("err3: %v\n", err3)

}
