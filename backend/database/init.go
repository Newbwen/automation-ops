package database

import (
	"automation-ops/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/automation_ops?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("无法连接数据库: ", err)
	}
	db.AutoMigrate(&models.Host{})
	DB = db
}
