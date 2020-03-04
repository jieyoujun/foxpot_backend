package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
	"github.com/likiiiiii/foxpot_backend/utils"
)

// DB 数据库
var (
	DB *gorm.DB
)

// InitDB 初始化数据库
func InitDB() (err error) {
	DB, err = gorm.Open(utils.GlobalConfig.DB.Type, utils.GlobalConfig.DB.DSN)
	if err != nil {
		return err
	}
	// 配置连接池
	DB.DB().SetMaxOpenConns(utils.GlobalConfig.DB.MaxOpenConn)
	DB.DB().SetMaxIdleConns(utils.GlobalConfig.DB.MaxIdleConn)
	DB.DB().SetConnMaxLifetime(time.Duration(utils.GlobalConfig.DB.MaxLifeTime) * time.Second)
	if !DB.HasTable(&User{}) {
		DB.AutoMigrate(&User{})
	}
	return nil
}
