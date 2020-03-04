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
	c.HTML(http.StatusOK, "user/index.html", gin.H{
		// 这里大家都有
		"code":    200,
		"message": "欢迎欢迎",
		"title":   "首页",
		// 这里是她小灶
		"User": user,
	})
}
