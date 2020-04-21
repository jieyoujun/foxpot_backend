package views

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetAdminIndex 管理员首页 默认攻击地图
func GetAdminIndex(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/attackmap", gin.H{
		"title": "攻击地图",
		"user":  user,
	})
}

// GetAdminProfile 个人中心
func GetAdminProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/profile", gin.H{
		"title": "个人中心",
		"user":  user,
	})
}

// GetAdminUpdateProfile 修改个人资料
func GetAdminUpdateProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/updateprofile", gin.H{
		"title": "修改个人资料",
		"user":  user,
	})
}

// PostAdminUpdateProfile 提交个人资料
func PostAdminUpdateProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	// TODO
	// 管理员用户修改个人资料也就更改密码 邮箱 手机号得了
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "admin/updateprofile", gin.H{
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
	if email != "" {
		user.Email = email
	}
	if phone != "" {
		user.Phone = phone
	}
	err := models.UpdateUser(user)
	if err != nil {
		c.HTML(http.StatusOK, "admin/updateprofile", gin.H{
			"message": "更新失败",
			"title":   "修改个人资料",
			"user":    user,
		})
		return
	}
	c.HTML(http.StatusOK, "admin/updateprofile", gin.H{
		"message": "更新成功",
		"title":   "修改个人资料",
		"user":    user,
	})
}

// GetAdminUserManage 用户管理
func GetAdminUserManage(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	users, _ := models.GetAllUsers()
	c.HTML(http.StatusOK, "admin/usermanage", gin.H{
		"title": "用户管理",
		"user":  user,
		"users": users,
	})
}

// GetAdminCreateUser 新建用户
func GetAdminCreateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/createuser", gin.H{
		"title": "新建用户",
		"user":  user,
	})
}

// PostAdminCreateUser 新建用户
func PostAdminCreateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm")
	role := models.Role2Uint(c.PostForm("role"))
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "admin/createuser", gin.H{
			"message": "用户名/密码不能为空",
			"title":   "新建用户",
			"user":    user,
		})
		return
	}
	// check if password != confirmPassword
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "admin/createuser", gin.H{
			"message": "密码不一致",
			"title":   "新建用户",
			"user":    user,
		})
		return
	}
	// check if superadmin
	if user.Role != 0 && role == 1 {
		c.HTML(http.StatusBadRequest, "admin/createuser", gin.H{
			"message": "只有超级管理员能够新建管理员",
			"title":   "新建用户",
			"user":    user,
		})
		return
	}
	// check if user already existed
	if _, err := models.GetUserByUsername(username); err == nil {
		c.HTML(http.StatusBadRequest, "admin/createuser", gin.H{
			"message": "用户名已存在",
			"title":   "新建用户",
			"user":    user,
		})
		return
	}
	password, _ = utils.HashPassword(password)
	fmt.Println(username, role)
	err := models.CreateUser(&models.User{
		Username:       username,
		HashedPassword: password,
		Role:           role,
		Email:          email,
		Phone:          phone,
	})
	if err != nil {
		c.HTML(http.StatusOK, "admin/createuser", gin.H{
			"message": "新建失败",
			"title":   "新建用户",
			"user":    user,
		})
	} else {
		c.HTML(http.StatusOK, "admin/createuser", gin.H{
			"message": "新建成功",
			"title":   "新建用户",
			"user":    user,
		})
	}
}

// PostAdminDeleteUser 删除用户
func PostAdminDeleteUser(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	var (
		usernames []string
		deleted   []string
	)
	c.ShouldBind(&usernames)
	if user.Role != 0 && models.HasAnyAdminByUsernames(usernames) {
		c.JSON(http.StatusOK, gin.H{
			"message": "只有超级管理员能够删除管理员",
		})
		return
	}
	for _, username := range usernames {
		// TODO 做成事务  错误回滚
		if err := models.DeleteUserByUsername(username); err == nil {
			deleted = append(deleted, username)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"data":    deleted,
	})
}

// GetAdminUpdateUser 更新用户信息
func GetAdminUpdateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	// TODO
	// 检查参数防乱搞
	username := c.Param("username")
	editUser, _ := models.GetUserByUsername(username)
	if user.Role != 0 && editUser.IsAdmin() {
		users, _ := models.GetAllUsers()
		c.HTML(http.StatusOK, "admin/usermanage", gin.H{
			"message": "只有超级管理员能够编辑管理员",
			"title":   "用户管理",
			"user":    user,
			"users":   users,
		})
		return
	}
	c.HTML(http.StatusOK, "admin/updateuser", gin.H{
		"title":    "编辑用户信息",
		"user":     user,
		"editUser": editUser,
	})
}

// PostAdminUpdateUser 更新用户信息
func PostAdminUpdateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	// TODO
	// 检查参数防乱搞
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm")
	role := models.Role2Uint(c.PostForm("role"))
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	editUser, _ := models.GetUserByUsername(username)
	// check if password != confirmPassword
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "admin/updateuser", gin.H{
			"message":  "密码不一致",
			"title":    "编辑用户信息",
			"user":     user,
			"editUser": editUser,
		})
		return
	}
	// check if superadmin
	if user.Role != 0 && role < editUser.Role {
		c.HTML(http.StatusBadRequest, "admin/updateuser", gin.H{
			"message":  "只有超级管理员能够提升权限",
			"title":    "编辑用户信息",
			"user":     user,
			"editUser": editUser,
		})
		return
	}
	if password != "" {
		password, _ = utils.HashPassword(password)
		editUser.HashedPassword = password
	}
	editUser.Role = role
	if email != "" {
		editUser.Email = email
	}
	if phone != "" {
		editUser.Phone = phone
	}
	err := models.UpdateUser(editUser)
	if err != nil {
		c.HTML(http.StatusOK, "admin/updateuser", gin.H{
			"message":  "更新失败",
			"title":    "编辑用户信息",
			"user":     user,
			"editUser": editUser,
		})
		return
	}
	c.HTML(http.StatusOK, "admin/updateuser", gin.H{
		"message":  "更新成功",
		"title":    "编辑用户信息",
		"user":     user,
		"editUser": editUser,
	})
}

// 第三方组件

// GetAdminESHead ESHead
func GetAdminESHead(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/eshead", gin.H{
		"title": "ES Head",
		"user":  user,
	})
}

// GetAdminKibanaDiscover KibanaDiscover
func GetAdminKibanaDiscover(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/discover", gin.H{
		"title": "告警详情",
		"user":  user,
	})
}

// GetAdminKibanaDashboard KibanaDashboard
func GetAdminKibanaDashboard(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/dashboard", gin.H{
		"title": "仪表板",
		"user":  user,
	})
}

// GetAdminCockpitDocker Cockpit容器状态监控
func GetAdminCockpitDocker(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/docker", gin.H{
		"title": "系统状态",
		"user":  user,
	})
}

// GetAdminCockpitSystem Cockpit系统状态监控
func GetAdminCockpitSystem(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/system", gin.H{
		"title": "系统状态",
		"user":  user,
	})
}

// GetAdminCockpitTerminal CockpitWeb终端
func GetAdminCockpitTerminal(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/terminal", gin.H{
		"title": "Web终端",
		"user":  user,
	})
}
