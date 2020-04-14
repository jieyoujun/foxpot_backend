package views

import (
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// Handle404 无效请求地址
func Handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"message": "404 Not Found",
		"title":   "错误",
	})
}

// GetIndex 根
func GetIndex(c *gin.Context) {
	session := sessions.Default(c)
	if userID, ok := session.Get(utils.Config.Session.Key).(uint); ok {
		if user, err := models.GetUserByID(userID); err == nil {
			if !user.IsAdmin() {
				c.Redirect(http.StatusFound, "/user")
				return
			}
			c.Redirect(http.StatusFound, "/admin")
			return
		}
	}
	c.Redirect(http.StatusFound, "/login")
}

// GetLogin 登录GET
func GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	if userID, ok := session.Get(utils.Config.Session.Key).(uint); ok {
		if user, err := models.GetUserByID(userID); err == nil {
			if !user.IsAdmin() {
				c.Redirect(http.StatusFound, "/user")
				return
			}
			c.Redirect(http.StatusFound, "/admin")
			return
		}
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
	})
}

// GetLogout 退出GET
func GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

// PostLogin 登录POST
func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"message": "用户名/密码不能为空",
			"title":   "登录",
		})
		return
	}
	captchaCode := c.PostForm("captcha")
	if captchaID, ok := session.Get("captcha").(string); !ok || !captcha.VerifyString(captchaID, captchaCode) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"message": "验证码错误",
			"title":   "登录",
		})
		return
	}
	user, err := models.GetUserByUsername(username)
	if err != nil || !utils.ValidatePassword(user.HashedPassword, password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"message": "用户名/密码不正确",
			"title":   "登录",
		})
		return
	}

	user.LastLoginAt = time.Now()
	models.UpdateUser(user)
	session.Clear()
	session.Set(utils.Config.Session.Key, user.ID)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
