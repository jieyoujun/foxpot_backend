package utils

import (
	"github.com/gin-contrib/sessions/cookie"
)

// NewCookieSessions ...
func NewCookieSessions() cookie.Store {
	store := cookie.NewStore([]byte(Secret))
	// store.Options(sessions.Options{
	// 	Path:     "/",
	// 	Domain:   "localhost",
	// 	MaxAge:   int(5 * time.Minute),
	// 	Secure:   true,
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteLaxMode,
	// })
	return store
}
