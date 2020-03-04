package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/likiiiiii/foxpot_backend/utils"
	"go.uber.org/zap"
)

// Logger 日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// Process request
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		var msg string
		switch {
		case statusCode >= 500:
			msg = "Server error – failed to fulfil an apparently valid request."
		case statusCode >= 400:
			msg = "Client error – request contains bad syntax or cannot be fulfilled."
		case statusCode >= 300:
			msg = "Redirection – further action needs to be taken in order to complete the request."
		case statusCode >= 200:
			msg = "Successful – request was successfully received, understood, and accepted."
		case statusCode >= 100:
			msg = "Informational response – request was received, continuing process."
		}
		if statusCode >= 400 {
			utils.Logger.Errorw(msg,
				zap.String("client_addr", c.Request.RemoteAddr),
				zap.String("method", c.Request.Method),
				zap.String("request_uri", c.Request.RequestURI),
				zap.String("proto", c.Request.Proto),
				zap.Int("status_code", statusCode),
				zap.String("ua", c.Request.UserAgent()),
				zap.String("latency", latency.String()),
			)
		} else {
			utils.Logger.Infow(msg,
				zap.String("client_addr", c.Request.RemoteAddr),
				zap.String("method", c.Request.Method),
				zap.String("request_uri", c.Request.RequestURI),
				zap.String("proto", c.Request.Proto),
				zap.Int("status_code", statusCode),
				zap.String("ua", c.Request.UserAgent()),
				zap.String("latency", latency.String()),
			)
		}
	}
}
