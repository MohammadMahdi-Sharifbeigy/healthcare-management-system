package main

import (
	"fmt"
	"log"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/config"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/database"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/http"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	log := logger.NewLogger(cfg.LogLevel, cfg.LogFormat)

	// Initialize database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Infof("Connected to database: %s:%d/%s", cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Info("Database migrations completed successfully")

	// Initialize HTTP server
	server := http.NewServer(cfg, db, log)

	// Start server
	address := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Infof("Starting Healthcare Management System API on %s (Environment: %s)", address, cfg.Environment)

	if err := server.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
