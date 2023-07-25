package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        c.Next()

        end := time.Now()

        latency := end.Sub(start)

        c.Set("logger", gin.H{
            "timestamp":   end,
            "statusCode":  c.Writer.Status(),
            "latency":     latency,
            "clientIP":    c.ClientIP(),
            "method":      c.Request.Method,
            "path":        c.Request.URL.Path,
            "userAgent":   c.Request.UserAgent(),
            "errors":      c.Errors.Errors(),
            "requestBody": c.Request.Body,
        })
    }
}
