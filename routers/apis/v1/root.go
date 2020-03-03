package apis

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetIndex ...
func GetIndex(c *gin.Context) {
	session := sessions.Default(c)
	if userID, ok := session.Get(utils.SessionKey).(uint); ok {
		if user, err := models.SelectUserByID(userID); err == nil {
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
	user, err := models.SelectUserByUsername(username)
	if err != nil || !utils.ComparePassword(user.HashedPassword, password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"code":    400,
			"message": "用户名或密码不正确",
		})
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.Set(utils.SessionKey, user.ID)
	session.Save()
	if user.IsAdmin() {
		c.Redirect(http.StatusMovedPermanently, "/admin")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/user")
}
