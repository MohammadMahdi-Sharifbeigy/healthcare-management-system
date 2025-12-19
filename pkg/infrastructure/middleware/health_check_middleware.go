package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthStatus struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    int64     `json:"uptime_seconds"`
	Version   string    `json:"version,omitempty"`
}

var serverStartTime time.Time

// InitHealth initializes health check
func InitHealth() {
	serverStartTime = time.Now()
}

// HealthCheckHandler returns server health status
func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uptime := int64(time.Since(serverStartTime).Seconds())

		status := HealthStatus{
			Status:    "healthy",
			Timestamp: time.Now(),
			Uptime:    uptime,
			Version:   "1.0.0",
		}

		c.JSON(http.StatusOK, status)
	}
}

// LivenessProbeHandler checks if server is running
func LivenessProbeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	}
}

// ReadinessProbeHandler checks if server is ready to accept requests
func ReadinessProbeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Add database connection check
		// For now, just return ready
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	}
}
