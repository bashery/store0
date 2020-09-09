package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func setdb() {

	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/store?charset=utf8&parseTime=True&loc=Local")
	//db, err = gorm.Open("mysql", "root:123456@/store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		panic("failed to connect database")

	}
}
