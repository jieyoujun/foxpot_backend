package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAdminIndex 管理员首页
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"code":    200,
		"message": "欢迎欢迎",
		"title":   "首页",
	})
}
