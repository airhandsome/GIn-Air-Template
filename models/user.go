package models

import "gorm.io/gorm"

// User 定义了与数据库表users对应的模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}

// InitModels 初始化模型，确保数据库表存在
func InitModels(db *gorm.DB) {
	// 自动迁移模型
	db.AutoMigrate(&User{})
}
