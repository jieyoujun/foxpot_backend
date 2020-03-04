package views

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/models"
	"github.com/likiiiiii/foxpot_backend/utils"
)

// GetUserIndex ...
func GetUserIndex(c *gin.Context) {
	session := sessions.Default(c)
	userID, _ := session.Get(utils.GlobalConfig.Session.Key).(uint)
	user, _ := models.GetUserByID(userID)
	c.HTML(http.StatusOK, "user/index.html", gin.H{
		"code":    200,
		"message": "",
		"User":    user,
	})
}
