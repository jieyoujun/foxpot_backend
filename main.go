package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/routers"
	"github.com/likiiiiii/foxpot_backend/utils"
)

func main() {
	configFilePath := flag.String("c", "etc/my.ini", "Path to config file")
	flag.Parse()
	// 0. 加载配置
	err := utils.LoadConfigFile(*configFilePath)
	if err != nil {
		log.Fatalln("Failed to parse config file:", err)
	}
	fmt.Printf("%#v\n", utils.GlobalConfig)
	// 1. 初始化
	if err := models.InitDB(); err != nil {
		log.Fatalln("Failed to init database:", err)
	}

	defer func() {
		if err := models.DB.Close(); err != nil {
			log.Fatalln("Failed to close db connection:", err)
		}
	}()
	gEngine := routers.Init()
	// 2. 启动
	if err := gEngine.Run(utils.GlobalConfig.Foxpot.Address); err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

// 插入测试用户
// hashedPassword, _ := utils.HashPassword("1212")
// models.CreateUser(&models.User{
// 	Username:       "liki",
// 	HashedPassword: hashedPassword,
// 	Role:           "admin",
// 	Email:          "admin@foxpot.com",
// 	Phone:          "19801209704",
// })
// models.CreateUser(&models.User{
// 	Username:       "niki",
// 	HashedPassword: hashedPassword,
// 	Role:           "user",
// 	Email:          "user@foxpot.com",
// 	Phone:          "19801209704",
// })
