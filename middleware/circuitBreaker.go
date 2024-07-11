package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CircuitBreakerMiddleware 熔断中间件
func CircuitBreakerMiddleware(failureThreshold int, recoveryTimeout time.Duration) gin.HandlerFunc {
	var failures int
	var lastFailureTime time.Time

	return func(c *gin.Context) {
		if failures >= failureThreshold && time.Since(lastFailureTime) < recoveryTimeout {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable due to circuit breaker"})
			c.Abort()
			return
		}

		c.Next()

		if c.Writer.Status() >= http.StatusInternalServerError {
			failures++
			lastFailureTime = time.Now()
		} else {
			failures = 0
		}
	}
}
