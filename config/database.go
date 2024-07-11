package config

import (
	"gin-air-template/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DatabaseConfig 包含了数据库配置信息
type DatabaseConfig struct {
	Dialect     string
	DSN         string
	Migrate     bool
	AutoMigrate bool
}

// InitDatabase 初始化数据库连接
func InitDatabase(dbConfig DatabaseConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if dbConfig.Dialect == "sqlite" {
		db, err = gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
	} else {
		// 添加其他数据库驱动的支持
	}

	if err != nil {
		return nil, err
	}

	if dbConfig.AutoMigrate {
		db.AutoMigrate(&models.User{}) // 自动迁移模型
	}

	return db, nil
}
