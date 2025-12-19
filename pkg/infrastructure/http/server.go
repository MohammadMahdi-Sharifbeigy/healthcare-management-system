package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/config"
)

type Server struct {
	router *gin.Engine
	config *config.Config
	db     *sql.DB
}

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func NewServer(cfg *config.Config, db *sql.DB, log Logger) *Server {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &Server{
		router: router,
		config: cfg,
		db:     db,
	}
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}
