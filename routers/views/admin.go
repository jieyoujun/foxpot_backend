package views

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetAdminIndex 管理员首页
func GetAdminIndex(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"code":    200,
		"message": "欢迎欢迎",
		"title":   "首页",
		"User":    user,
	})
}

// GetAdminUserManage 用户管理
func GetAdminUserManage(c *gin.Context) {
	users, _ := models.GetAllUsers()
	fmt.Println(users)
	c.HTML(http.StatusOK, "admin/usermanage.html", gin.H{
		"code":    200,
		"message": "用户列表",
		"title":   "用户管理",
		"Users":   users,
	})
}
