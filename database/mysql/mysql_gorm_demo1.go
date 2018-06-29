package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

type Userinfo struct {
	Uid int
	Username string
	Departname string
	Created time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:789456@/golang?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("连接数据库失败")
	}
	exists := db.HasTable(&Userinfo{})
	fmt.Println(exists)

	fmt.Println("连接数据库成功")
	db.AutoMigrate(&Userinfo{})
	user := Userinfo{Username:"Tom",Departname:"dep",Created:time.Now()}
	fmt.Println(user)
	db.Create(&user)
	fmt.Println(db.First(&user))
}