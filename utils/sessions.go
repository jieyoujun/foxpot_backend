package utils

import (
	"github.com/gin-contrib/sessions/cookie"
)

// TODO
// 这块感觉缺点啥

// NewCookieSessions 创建session库
func NewCookieSessions(secret string) cookie.Store {
	store := cookie.NewStore([]byte(secret))
	return store
}
