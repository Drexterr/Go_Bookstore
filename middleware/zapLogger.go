package middleware

import (
	"time"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		initializers.Log.Info("Request",
			zap.String("method", c.Request.Method),
			zap.Int("Status", c.Writer.Status()),
			zap.String("Path", path),
			zap.String("Query", query),
			zap.String("ClientIP", c.ClientIP()),
			zap.String("UserAgent", c.Request.UserAgent()),
			zap.Duration("Duration", time.Since(start)),
			zap.String("ErrorMessage", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}

}
