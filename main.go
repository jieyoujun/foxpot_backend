package main

import (
	"flag"
	"log"

	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/routers"
	"github.com/likiiiiii/foxpot_backend/utils"
	"go.uber.org/zap"
)

// 把大象塞进冰箱分几步
// 答曰: 三步
// 1.把冰箱门打开
// 2.把大象塞进去
// 3.把冰箱门带上

// 创建后台分几步
// 答曰: 同上
// 0.加载配置
// 1.初始化
// 2.跑

func main() {
	configFilePath := flag.String("c", "etc/my.ini", "Path to config file")
	flag.Parse()
	// 0. 加载配置
	if err := utils.LoadConfigFile(*configFilePath); err != nil {
		log.Fatalf("Failed to load config from %v, err: %v\n", *configFilePath, err)
	}
	// 1. 初始化
	// 1.1 日志
	if err := utils.InitLogger(); err != nil {
		log.Fatalf("Failed to initiate logger, err: %v\n", err)
	}
	// >>>>>>>>>>>>>>>>>>>>日志起于此处<<<<<<<<<<<<<<<<<<<<
	// 1.2 数据库
	if err := models.InitDB(); err != nil {
		utils.Logger.Fatalw("Failed to initiate database.", zap.String("Error", err.Error()))
	}
	utils.Logger.Debugw("Succeed to initiate database.")
	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// 需要安全退出的请在1.3之前初始化
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	// 1.3 安全退出
	defer func() {
		if err := models.DB.Close(); err != nil {
			utils.Logger.Fatalw("Failed to close database clearly.", zap.String("Error", err.Error()))
		}
		if err := utils.Logger.Sync(); err != nil {
			log.Fatalf("Failed to close logger clearly, err: %v\n", err)
		}
	}()
	// 1.4 HTTP服务
	routers.Init()
	utils.Logger.Debugw("Succeed to initiate router.")
	// 2. 启动
	if err := routers.GEngine.Run(utils.Config.Foxpot.Address); err != nil {
		utils.Logger.Fatalw("Failed to start server.", zap.String("Error", err.Error()))
	}
}
