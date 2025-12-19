package logger

import (
	"log"
	"os"
)

// Logger provides structured logging functionality
type Logger struct {
	level  string
	format string
	*log.Logger
}

// NewLogger creates a new logger instance
func NewLogger(level, format string) *Logger {
	return &Logger{
		level:  level,
		format: format,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Info logs an info level message
func (l *Logger) Info(msg string) {
	l.Printf("[INFO] %s", msg)
}

// Infof logs an info level message with format
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Printf("[INFO] "+format, args...)
}

// Debug logs a debug level message
func (l *Logger) Debug(msg string) {
	if l.level == "debug" {
		l.Printf("[DEBUG] %s", msg)
	}
}

// Debugf logs a debug level message with format
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level == "debug" {
		l.Printf("[DEBUG] "+format, args...)
	}
}

// Warn logs a warning level message
func (l *Logger) Warn(msg string) {
	l.Printf("[WARN] %s", msg)
}

// Warnf logs a warning level message with format
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Printf("[WARN] "+format, args...)
}

// Error logs an error level message
func (l *Logger) Error(msg string) {
	l.Printf("[ERROR] %s", msg)
}

// Errorf logs an error level message with format
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Printf("[ERROR] "+format, args...)
}

// Fatal logs a fatal level message and exits
func (l *Logger) Fatal(msg string) {
	l.Fatalf("[FATAL] %s", msg)
}

// Fatalf logs a fatal level message with format and exits
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf("[FATAL] "+format, args...)
}
