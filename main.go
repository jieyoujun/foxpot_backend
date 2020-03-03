package main

import (
	"log"

	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/routers"
	"github.com/likiiiiii/foxpot_backend/utils"
)

var ()

func main() {
	// 0. 加载配置
	utils.LoadConfig()
	// 1. 初始化
	models.InitDBAndMigrateAll()
	defer models.DB.Close()
	gEngine := routers.Init()
	// 2. 启动
	if err := gEngine.Run(); err != nil {
		log.Fatalln("Failed to start server")
	}
}

// 插入测试用户
// hashedPassword, _ := utils.HashPassword("1212")
// models.InsertUser(&models.User{
// 	Username:       "liki",
// 	Password:       "1212",
// 	HashedPassword: hashedPassword,
// 	Role:           "administrator",
// 	Email:          "liki@foxpot.com",
// 	Phone:          "19801209704",
// })
// models.InsertUser(&models.User{
// 	Username:       "niki",
// 	Password:       "1212",
// 	HashedPassword: hashedPassword,
// 	Role:           "user",
// 	Email:          "niki@foxpot.com",
// 	Phone:          "19801209704",
// })
