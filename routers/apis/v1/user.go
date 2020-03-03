package apis

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
	userID, _ := session.Get(utils.SessionKey).(uint)
	user, _ := models.SelectUserByID(userID)
	c.HTML(http.StatusOK, "user/index.html", gin.H{
		"code":    200,
		"message": "",
		"User":    user,
	})
}
