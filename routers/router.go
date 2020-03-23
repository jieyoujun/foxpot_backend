package routers

import (
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	v1 "github.com/likiiiiii/foxpot_backend/routers/apis/v1"
	"github.com/likiiiiii/foxpot_backend/routers/middlewares"
	"github.com/likiiiiii/foxpot_backend/routers/views"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GEngine HTTP服务
var GEngine *gin.Engine

// Init 初始化gin引擎
func Init() {
	GEngine = gin.New()
	// 1.4.1 注册模板函数
	// >>>>>>>>>>>>>>>>>
	// 这里写注册模板函数
	GEngine.SetFuncMap(template.FuncMap{
		"autoincrement": func(num int) int {
			return num + 1
		},
	})
	// <<<<<<<<<<<<<<<<<
	// 1.4.2 注册模板文件
	GEngine.LoadHTMLGlob("views/**/*")
	// 1.4.3 注册静态文件
	GEngine.Static("/statics", "./statics")
	// 1.4.4 注册中间件、路由
	// >>>>>>>>>>>
	// 这里是中间件
	// <<<<<<<<<<<
	GEngine.Use(gin.Logger(), gin.Recovery())
	GEngine.Use(middlewares.Logger(), sessions.Sessions(utils.Config.Session.Key, utils.NewCookieSessions(utils.Config.Session.Secret)))
	// >>>>>>>>>
	// 这里是路由
	// <<<<<<<<<
	GEngine.GET("/", views.GetIndex)
	GEngine.GET("/login", views.GetLogin)
	GEngine.GET("/logout", views.GetLogout)
	GEngine.POST("/login", views.PostLogin)
	// >>>>>>>>>>>
	// 这里是路由组
	// <<<<<<<<<<<
	admin := GEngine.Group("/admin", middlewares.AdminRequired())
	{
		admin.GET("/", views.GetAdminIndex)
		admin.GET("/usermanage", views.GetAdminUserManage)
	}
	user := GEngine.Group("/user", middlewares.UserRequired())
	{
		user.GET("/", views.GetUserIndex)
	}

	apiv1 := GEngine.Group("/api/v1")
	{
		apiv1.GET("/captcha", v1.GetCaptcha)
	}
	// >>>>>>>>>
	// 这里是404
	// <<<<<<<<<
	GEngine.NoRoute(views.Handle404)
}
