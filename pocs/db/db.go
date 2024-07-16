package db

import (
	"log"

	"demo/global"
	"demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "root:Snowfish3.14@tcp(127.0.0.1:3306)/pocs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 自动迁移模式，创建表
	if err := db.AutoMigrate(&model.FileRecord{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	global.DB = db
	return db
}
