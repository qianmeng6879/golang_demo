package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

var dsn string = "root:root@tcp(127.0.0.1:3306)/golang_demo?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败！")
		fmt.Println(err)
	}

	Db = db
}
