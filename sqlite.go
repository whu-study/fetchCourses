package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func init() {
	var err error
	var cfg gorm.Config
	cfg = gorm.Config{
		PrepareStmt: true,
		ConnPool:    nil,
	}
	// 连接到SQLite数据库
	if Client, err = gorm.Open(sqlite.Open("resources/test.db"), &cfg); err != nil {
		panic(err)
	}

	TableAutoMigrate()
}

func TableAutoMigrate() {
	if err := Client.AutoMigrate(&TimeInfo{}, &CourseInfo{}); err != nil {
		panic(err)
		return
	}

}
