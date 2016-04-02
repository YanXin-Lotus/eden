package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	DB gorm.DB
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open("postgres", "user=postgres password=root dbname=eden sslmode=disable")
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
		fmt.Println("创建 User 表失败！")
	}

	err = DB.CreateTable(&Article{}).Error
	if err != nil {
		fmt.Println("创建 Article 表失败！")
	}
}
