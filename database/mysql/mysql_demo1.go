package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	"time"
	_"github.com/ziutek/mymysql/godrv"
)

type Userinfo struct {
	Uid int
	Username string
	Departname string
	Created time.Time
}


func main() {
	//db, err := sql.Open("mysql", "root:789456@/golang?charset=utf8")
	db, err := sql.Open("mymysql","golang/root/789456")
	defer db.Close()
	if err != nil{
		//panic(err)
		fmt.Println(22222)
	}
	fmt.Println("连接数据库成功")

	beedb.OnDebug = true
	orm := beedb.New(db)

	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)
}