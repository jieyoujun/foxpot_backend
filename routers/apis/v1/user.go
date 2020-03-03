package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserIndex ...
func GetUserIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.html", gin.H{
		"title": "User",
	})
}
