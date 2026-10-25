package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		Log.Info(
			"http request",
			zap.String(
				"method",
				c.Request.Method,
			),
			zap.String(
				"path",
				c.Request.URL.Path,
			),
			zap.Int(
				"status",
				c.Writer.Status(),
			),
			zap.Duration(
				"latency",
				time.Since(start),
			),
		)
	}
}
