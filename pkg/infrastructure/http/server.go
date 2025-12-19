package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Port                 int
	Environment          string
	AllowedOrigins       []string
	MaxRequestSize       int64
	RateLimitEnabled     bool
	RateLimitRequests    int
	RateLimitWindow      time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	ShutdownTimeout      time.Duration
}

type Server struct {
	router *gin.Engine
	srv    *http.Server
	config *RouterConfig
}

// NewServer creates a new HTTP server
func NewServer(config *RouterConfig, handlers *Handlers) *Server {
	// Set Gin mode
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// Setup routes with middleware
	SetupRoutes(router, handlers, config)

	// Create HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	return &Server{
		router: router,
		srv:    srv,
		config: config,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	fmt.Printf("🚀 Starting HTTP server on %s (Environment: %s)\n", s.srv.Addr, s.config.Environment)
	fmt.Printf("📊 Rate limiting: %v\n", s.config.RateLimitEnabled)
	fmt.Printf("🔒 CORS origins: %v\n", s.config.AllowedOrigins)

	return s.srv.ListenAndServe()
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("🛑 Shutting down HTTP server...")

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(ctx, s.config.ShutdownTimeout)
	defer cancel()

	return s.srv.Shutdown(shutdownCtx)
}

// GetRouter returns the Gin router
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

// GetAddr returns the server address
func (s *Server) GetAddr() string {
	return s.srv.Addr
}
