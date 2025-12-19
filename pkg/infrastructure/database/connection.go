package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Database        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type Database struct {
	conn *sql.DB
}

// NewDatabase creates new database connection
func NewDatabase(config *DatabaseConfig) (*Database, error) {
	dsn := buildDSN(config)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return &Database{conn: db}, nil
}

// buildDSN constructs PostgreSQL connection string
func buildDSN(config *DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)
}

// GetConnection returns raw sql.DB connection
func (d *Database) GetConnection() *sql.DB {
	return d.conn
}

// Close closes database connection
func (d *Database) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
	return nil
}

// Ping tests database connectivity
func (d *Database) Ping() error {
	return d.conn.Ping()
}

// HealthCheck performs comprehensive health check
func (d *Database) HealthCheck() error {
	// Test basic connectivity
	if err := d.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	// Get connection pool stats
	stats := d.conn.Stats()

	// Check for connection pool exhaustion
	if stats.OpenConnections >= stats.MaxOpenConnections {
		return fmt.Errorf("connection pool exhausted: %d/%d", stats.OpenConnections, stats.MaxOpenConnections)
	}

	return nil
}

// ConnectionStats returns database connection statistics
func (d *Database) ConnectionStats() sql.DBStats {
	return d.conn.Stats()
}

// DefaultConfig returns default database configuration
func DefaultConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:            "localhost",
		Port:            "5432",
		User:            "postgres",
		Password:        "postgres",
		Database:        "healthcare_db",
		SSLMode:         "disable",
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	}
}

// ProductionConfig returns production database configuration
func ProductionConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:            "prod-db.example.com",
		Port:            "5432",
		User:            "healthcare_user",
		Password:        "", // Load from environment
		Database:        "healthcare_db",
		SSLMode:         "require",
		MaxOpenConns:    50,
		MaxIdleConns:    10,
		ConnMaxLifetime: 15 * time.Minute,
		ConnMaxIdleTime: 5 * time.Minute,
	}
}

// DevelopmentConfig returns development database configuration
func DevelopmentConfig() *DatabaseConfig {
	return DefaultConfig()
}

// TestConfig returns test database configuration
func TestConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:            "localhost",
		Port:            "5432",
		User:            "postgres",
		Password:        "postgres",
		Database:        "healthcare_test_db",
		SSLMode:         "disable",
		MaxOpenConns:    10,
		MaxIdleConns:    2,
		ConnMaxLifetime: 1 * time.Minute,
		ConnMaxIdleTime: 30 * time.Second,
	}
}

// DatabaseInfo returns database information
type DatabaseInfo struct {
	Version      string
	Database     string
	User         string
	CurrentTime  time.Time
	IsConnected  bool
	ConnectionStats sql.DBStats
}

// GetDatabaseInfo retrieves database information
func (d *Database) GetDatabaseInfo() (*DatabaseInfo, error) {
	var version, database, user string
	var currentTime time.Time

	// Get PostgreSQL version
	err := d.conn.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return nil, err
	}

	// Get current database name
	err = d.conn.QueryRow("SELECT current_database()").Scan(&database)
	if err != nil {
		return nil, err
	}

	// Get current user
	err = d.conn.QueryRow("SELECT current_user").Scan(&user)
	if err != nil {
		return nil, err
	}

	// Get current time
	err = d.conn.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		return nil, err
	}

	return &DatabaseInfo{
		Version:         version,
		Database:        database,
		User:            user,
		CurrentTime:     currentTime,
		IsConnected:     d.Ping() == nil,
		ConnectionStats: d.ConnectionStats(),
	}, nil
}
