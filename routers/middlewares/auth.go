package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// UserRequired ...
func UserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if userID, ok := session.Get(utils.GlobalConfig.Session.Key).(uint); ok {
			if user, err := models.GetUserByID(userID); err == nil && !user.IsAdmin() {
				c.Next()
				return
			}
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"code":    403,
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Redirect(http.StatusFound, "/login")
	}
}

// AdminRequired ...
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if userID, ok := session.Get(utils.GlobalConfig.Session.Key).(uint); ok {
			if user, err := models.GetUserByID(userID); err == nil && user.IsAdmin() {
				c.Next()
				return
			}
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"code":    403,
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Redirect(http.StatusFound, "/login")
	}
}
