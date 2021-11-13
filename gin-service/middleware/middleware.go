package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"omar":  "123",
		"barra": "678",
	})
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
