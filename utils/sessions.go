package utils

import (
	"github.com/gin-contrib/sessions/cookie"
)

// NewCookieSessions ...
func NewCookieSessions(secret string) cookie.Store {
	store := cookie.NewStore([]byte(secret))
	return store
}
