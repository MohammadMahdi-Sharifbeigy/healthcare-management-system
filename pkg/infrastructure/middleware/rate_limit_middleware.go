package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
// limit: max requests per window
// window: time window duration
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Cleanup old entries every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			rl.cleanup()
		}
	}()

	return rl
}

// IsAllowed checks if request should be allowed
func (rl *RateLimiter) IsAllowed(clientIP string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	timestamps := rl.requests[clientIP]

	// Remove old timestamps outside window
	validTimestamps := []time.Time{}
	for _, ts := range timestamps {
		if now.Sub(ts) < rl.window {
			validTimestamps = append(validTimestamps, ts)
		}
	}

	// Check if limit exceeded
	if len(validTimestamps) >= rl.limit {
		return false
	}

	// Add current request
	validTimestamps = append(validTimestamps, now)
	rl.requests[clientIP] = validTimestamps

	return true
}

// cleanup removes stale entries
func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	for ip, timestamps := range rl.requests {
		validTimestamps := []time.Time{}
		for _, ts := range timestamps {
			if now.Sub(ts) < rl.window {
				validTimestamps = append(validTimestamps, ts)
			}
		}

		if len(validTimestamps) == 0 {
			delete(rl.requests, ip)
		} else {
			rl.requests[ip] = validTimestamps
		}
	}
}

// RateLimitMiddleware returns gin middleware for rate limiting
// limit: requests per window
// window: duration (e.g., 1*time.Minute for 1 minute window)
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(limit, window)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !limiter.IsAllowed(clientIP) {
			c.Header("Retry-After", "60")
			c.JSON(http.StatusTooManyRequests, ErrorResponse{
				Error:   "RATE_LIMIT_EXCEEDED",
				Message: "Too many requests. Please try again later.",
				Code:    http.StatusTooManyRequests,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
