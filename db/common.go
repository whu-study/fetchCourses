package database

import (
	"fetchCourses/config"
	"fmt"
	"gorm.io/gorm"
)

var Client *gorm.DB
var isInit bool

func TableAutoMigrate() {
	if !config.Conf.Mysql.AutoMigrate {
		fmt.Println("未启用迁移数据库")
		return
	}

	// 批量执行自动迁移
	if err := Client.AutoMigrate(&CourseInfo{}, &TimeInfo{}); err != nil {
		panic(fmt.Errorf("数据库迁移失败: %v", err))
	}
}
