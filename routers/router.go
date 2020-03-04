package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/routers/middlewares"
	"github.com/likiiiiii/foxpot_backend/routers/views"
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
	gEngine.Use(sessions.Sessions(utils.GlobalConfig.Session.Key, utils.NewCookieSessions(utils.GlobalConfig.Session.Secret)))
	gEngine.GET("/", views.GetIndex)
	gEngine.GET("/login", views.GetLogin)
	gEngine.GET("/logout", views.GetLogout)
	gEngine.POST("/login", views.PostLogin)
	user := gEngine.Group("/user", middlewares.UserRequired())
	{
		user.GET("/", views.GetUserIndex)
	}
	admin := gEngine.Group("/admin", middlewares.AdminRequired())
	{
		admin.GET("/", views.GetAdminIndex)
	}
	gEngine.NoRoute(views.Handle404)
	// 5. 完事儿
	return gEngine
}
