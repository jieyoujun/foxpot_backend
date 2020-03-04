package views

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// Handle404 ...
func Handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"code":    404,
		"message": "宁迷路了",
	})
}

// GetIndex ...
func GetIndex(c *gin.Context) {
	session := sessions.Default(c)
	if userID, ok := session.Get(utils.GlobalConfig.Session.Key).(uint); ok {
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

// GetLogin ...
func GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	if userID, ok := session.Get(utils.GlobalConfig.Session.Key).(uint); ok {
		if user, err := models.GetUserByID(userID); err == nil {
			if !user.IsAdmin() {
				c.Redirect(http.StatusFound, "/user")
				return
			}
			c.Redirect(http.StatusFound, "/admin")
			return
		}
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

// GetLogout ...
func GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

// PostLogin ...
func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"code":    400,
			"message": "用户名和密码不能为空",
		})
		return
	}
	user, err := models.GetUserByUsername(username)
	if err != nil || !utils.ValidatePassword(user.HashedPassword, password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"code":    400,
			"message": "用户名或密码不正确",
		})
		return
	}
	user.LastLoginAt = time.Now()
	models.UpdateUser(user)
	session := sessions.Default(c)
	session.Clear()
	session.Set(utils.GlobalConfig.Session.Key, user.ID)
	session.Save()
	if user.IsAdmin() {
		c.Redirect(http.StatusMovedPermanently, "/admin")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/user")
}
