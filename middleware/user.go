package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 相关头部信息
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")         // 允许所有来源（根据需要限制特定来源）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 是否允许发送 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 对 OPTIONS 请求直接返回
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 继续后续的处理
		c.Next()
	}
}
