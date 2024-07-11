package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiterMiddleware 限流中间件
func RateLimiterMiddleware(limiter *rate.Limiter) gin.HandlerFunc {
	var mu sync.Mutex
	visitors := make(map[string]*rate.Limiter)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		mu.Lock()
		limiter, exists := visitors[clientIP]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Minute), 10) // 每分钟最多10次请求
			visitors[clientIP] = limiter
		}
		mu.Unlock()

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
