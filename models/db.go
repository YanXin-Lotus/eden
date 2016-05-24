package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@/eden?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	DB.DB()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.SingularTable(true)

	err = DB.CreateTable(&User{}).Error
	if err != nil {
		fmt.Println("创建 User 表失败(请检查User表是否已经存在)！")
	}

	err = DB.CreateTable(&Article{}).Error
	if err != nil {
		fmt.Println("创建 Article 表失败(请检查Article表是否已经存在)！")
	}
    
    err = DB.CreateTable(&Note{}).Error
	if err != nil {
		fmt.Println("创建 Note 表失败(请检查Note表是否已经存在)！")
	}
}
