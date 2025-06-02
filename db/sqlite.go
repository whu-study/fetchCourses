package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() {
	if isInit {
		return
	}
	isInit = true

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
