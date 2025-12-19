package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	ClientIP    string    `json:"client_ip"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	StatusCode  int       `json:"status_code"`
	ResponseTime int64    `json:"response_time_ms"`
	UserAgent   string    `json:"user_agent"`
	RequestID   string    `json:"request_id"`
}

// LoggerMiddleware logs all incoming HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		// Generate unique request ID
		requestID := fmt.Sprintf("%d-%s", startTime.UnixNano(), c.ClientIP())
		c.Set("RequestID", requestID)

		// Log request details
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Process request
		c.Next()

		// Calculate response time
		responseTime := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()

		// Log response details
		logEntry := LogEntry{
			Timestamp:   startTime,
			ClientIP:    clientIP,
			Method:      method,
			Path:        path,
			StatusCode:  statusCode,
			ResponseTime: responseTime,
			UserAgent:   userAgent,
			RequestID:   requestID,
		}

		// Log based on status code
		logMessage := fmt.Sprintf(
			"[%s] %s %s - Status: %d - Duration: %dms - IP: %s",
			logEntry.RequestID,
			method,
			path,
			statusCode,
			responseTime,
			clientIP,
		)

		if statusCode >= 500 {
			fmt.Printf("[ERROR] %s\n", logMessage)
		} else if statusCode >= 400 {
			fmt.Printf("[WARN] %s\n", logMessage)
		} else {
			fmt.Printf("[INFO] %s\n", logMessage)
		}

		// Log slow requests (>1s)
		if responseTime > 1000 {
			fmt.Printf("[SLOW] %s (Duration: %dms)\n", logMessage, responseTime)
		}
	}
}

// RequestIDMiddleware extracts or generates request ID
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID := fmt.Sprintf("%d", time.Now().UnixNano())
			c.Set("RequestID", requestID)
		} else {
			c.Set("RequestID", requestID)
		}
		
		// Add request ID to response header
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
