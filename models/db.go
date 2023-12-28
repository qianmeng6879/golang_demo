package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

var dsn string = "root:root@tcp(127.0.0.1:3306)/golang_demo?charset=utf8mb4&parseTime=True&loc=Local"
var dsn2 string = "root:root@tcp(192.168.2.150:3306)/golang_demo?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("1-数据库连接失败！")
		fmt.Println(err)
		db, err = gorm.Open(mysql.Open(dsn2), &gorm.Config{})
		if err != nil {
			fmt.Println("2-数据库连接失败！")
			fmt.Println(err)
		}
	}
	Db = db
}
