package views

import (
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

// GetAdminProfile 管理员个人中心
func GetAdminProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.Config.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "admin/profile.html", gin.H{
		"code":    200,
		"message": "管理员个人中心",
		"title":   "个人中心",
		"User":    user,
	})
}

// GetAdminUserManage 用户管理
func GetAdminUserManage(c *gin.Context) {
	users, _ := models.GetAllUsers()
	c.HTML(http.StatusOK, "admin/usermanage.html", gin.H{
		"code":    200,
		"message": "用户列表",
		"title":   "用户管理",
		"Users":   users,
	})
}

// GetAdminCreateUser 新建用户
func GetAdminCreateUser(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/createuser.html", gin.H{
		"code":    200,
		"message": "",
		"title":   "新建用户",
	})
}

// PostAdminCreateUser 新建用户
func PostAdminCreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	if username == "" || password == "" || role == "" {
		c.HTML(http.StatusBadRequest, "admin/createuser.html", gin.H{
			"code":    400,
			"message": "用户名和密码不能为空",
			"title":   "新建用户",
		})
		return
	}
	// check if already existed
	if _, err := models.GetUserByUsername(username); err == nil {
		c.HTML(http.StatusBadRequest, "admin/createuser.html", gin.H{
			"code":    400,
			"message": "用户已存在",
			"title":   "新建用户",
		})
		return
	}
	password, _ = utils.HashPassword(password)
	err := models.CreateUser(&models.User{
		Username:       username,
		HashedPassword: password,
		Role:           role,
		Email:          email,
		Phone:          phone,
	})
	if err != nil {
		c.HTML(http.StatusOK, "admin/createuser.html", gin.H{
			"code":    500,
			"message": "新建失败",
			"title":   "新建用户",
		})
	} else {
		c.HTML(http.StatusOK, "admin/createuser.html", gin.H{
			"code":    200,
			"message": "新建成功",
			"title":   "新建用户",
		})
	}
}

// PostAdminDeleteUser 删除用户
func PostAdminDeleteUser(c *gin.Context) {
	var (
		usernames []string
		deleted   []string
	)
	c.ShouldBind(&usernames)
	for _, username := range usernames {
		// TODO 做成事务  错误回滚
		if err := models.DeleteUserByUsername(username); err == nil {
			deleted = append(deleted, username)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data":    deleted,
	})
}

// GetAdminUpdateUser 更新用户信息
func GetAdminUpdateUser(c *gin.Context) {
	// TODO 参数判定下 这样不严格 防乱搞
	username := c.DefaultQuery("username", "none")
	user, _ := models.GetUserByUsername(username)
	c.HTML(http.StatusOK, "admin/updateuser.html", gin.H{
		"code":    200,
		"message": "",
		"title":   "更新用户信息",
		"User":    user,
	})
}

// PostAdminUpdateUser 更新用户信息
func PostAdminUpdateUser(c *gin.Context) {
	username := c.PostForm("username")
	user, _ := models.GetUserByUsername(username)
	user.Role = c.PostForm("role")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	err := models.UpdateUser(user)
	if err != nil {
		c.HTML(http.StatusOK, "admin/updateuser.html", gin.H{
			"code":    500,
			"message": "更新失败",
			"title":   "更新用户信息",
			"User":    user,
		})
	} else {
		c.HTML(http.StatusOK, "admin/updateuser.html", gin.H{
			"code":    200,
			"message": "更新成功",
			"title":   "更新用户信息",
			"User":    user,
		})
	}
}
