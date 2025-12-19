package http

import (
	"time"
)

// DefaultConfig returns default router configuration
func DefaultConfig() *RouterConfig {
	return &RouterConfig{
		Port:            8080,
		Environment:     "development",
		AllowedOrigins:  []string{"http://localhost:3000", "http://localhost:5173"},
		MaxRequestSize:  10 << 20, // 10 MB
		RateLimitEnabled: true,
		RateLimitRequests: 100,
		RateLimitWindow:   1 * time.Minute,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ShutdownTimeout:   10 * time.Second,
	}
}

// ProductionConfig returns production configuration
func ProductionConfig() *RouterConfig {
	return &RouterConfig{
		Port:            80,
		Environment:     "production",
		AllowedOrigins:  []string{"https://hospital.example.com"},
		MaxRequestSize:  5 << 20, // 5 MB
		RateLimitEnabled: true,
		RateLimitRequests: 50,
		RateLimitWindow:   1 * time.Minute,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
		ShutdownTimeout:   30 * time.Second,
	}
}

// DevelopmentConfig returns development configuration
func DevelopmentConfig() *RouterConfig {
	return DefaultConfig()
}
