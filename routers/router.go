package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/routers/apis/v1"
	"github.com/likiiiiii/foxpot_backend/routers/middlewares"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// Init ...
func Init() *gin.Engine {
	// 4. 创建server
	gEngine := gin.Default()
	// 4.1 注册模板函数
	// 4.2 注册模板文件
	gEngine.LoadHTMLGlob("views/**/*")
	// 4.3 注册静态文件
	gEngine.Static("/statics", "./statics")
	// 4.4 注册中间件、路由
	gEngine.Use(sessions.Sessions(utils.SessionKey, utils.NewCookieSessions()))
	gEngine.GET("/", apis.GetIndex)
	gEngine.GET("/login", apis.GetLogin)
	gEngine.GET("/logout", apis.GetLogout)
	gEngine.POST("/login", apis.PostLogin)
	user := gEngine.Group("/user", middlewares.UserRequired())
	{
		user.GET("/", apis.GetUserIndex)
	}
	admin := gEngine.Group("/admin", middlewares.AdminRequired())
	{
		admin.GET("/", apis.GetAdminIndex)
	}
	// 5. 完事儿
	return gEngine
}
