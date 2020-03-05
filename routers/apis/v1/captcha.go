package v1

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// TODO
// 支持ES查询
// 攻击事件给到攻击地图
// ...

// GetCaptcha 验证码
func GetCaptcha(c *gin.Context) {
	session := sessions.Default(c)
	captchaID := captcha.New()
	session.Delete("captcha")
	session.Set("captcha", captchaID)
	session.Save()
	captcha.WriteImage(c.Writer, captchaID, captcha.StdWidth, captcha.StdHeight)
}
