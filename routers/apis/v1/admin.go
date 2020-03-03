package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAdminIndex ...
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"code":    200,
		"message": "",
	})
}
