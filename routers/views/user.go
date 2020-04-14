package views

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetUserIndex 用户首页
func GetUserIndex(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "user/attackmap", gin.H{
		"title": "攻击地图",
		"user":  user,
	})
}

// 个人中心

// GetUserProfile 个人中心
func GetUserProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "user/profile", gin.H{
		"title": "个人中心",
		"user":  user,
	})
}

// GetUserUpdateProfile 修改个人资料
func GetUserUpdateProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "user/updateprofile", gin.H{
		"title": "修改个人资料",
		"user":  user,
	})
}

// PostUserUpdateProfile 提交个人资料
func PostUserUpdateProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	// TODO
	// 普通用户修改个人资料就更改密码 邮箱 手机号得了
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "user/updateprofile", gin.H{
			"message": "密码不一致",
			"title":   "修改个人资料",
			"user":    user,
		})
		return
	}
	if password != "" {
		password, _ = utils.HashPassword(password)
		user.HashedPassword = password
	}
	user.Email = email
	user.Phone = phone
	err := models.UpdateUser(user)
	if err != nil {
		c.HTML(http.StatusOK, "user/updateprofile", gin.H{
			"message": "更新失败",
			"title":   "修改个人资料",
			"user":    user,
		})
	} else {
		c.HTML(http.StatusOK, "user/updateprofile", gin.H{
			"message": "更新成功",
			"title":   "修改个人资料",
			"user":    user,
		})
	}
}

// 第三方组件

// GetUserKibanaDiscover KibanaDiscover
func GetUserKibanaDiscover(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/discover", gin.H{
		"title": "告警详情",
		"user":  user,
	})
}

// GetUserKibanaDashboard KibanaDashboard
func GetUserKibanaDashboard(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/dashborad.html", gin.H{
		"title": "仪表板",
		"user":  user,
	})
}

// GetUserCockpitSystem Cockpit系统状态监控
func GetUserCockpitSystem(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/system", gin.H{
		"title": "系统状态",
		"User":  user,
	})
}
