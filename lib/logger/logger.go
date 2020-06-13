package logger

import (
	"log"

	"github.com/nicosrgh/straw-hat/config"
)

const (
	// LevelDebug :
	LevelDebug = "DEBUG"
	// LevelInfo  :
	LevelInfo = "INFO"
	// LevelError :
	LevelError = "ERROR"
)

var logger Logger

func init() {
	logger = NewStandardLogger()
	if config.C.LogLevel == "" {
		config.C.LogLevel = LevelDebug
	}
}

// Debug ...
func Debug(msg string, args ...interface{}) {
	if config.C.LogLevel == LevelDebug {
		logger.Debug(msg, args...)
	}
}

// Info ...
func Info(msg string, args ...interface{}) {
	if config.C.LogLevel == LevelInfo || config.C.LogLevel == LevelDebug {
		logger.Info(msg, args...)
	}
}

// Error ...
func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

// Logger ...
type Logger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// NewStandardLogger ...
func NewStandardLogger() *StandardLogger {
	return new(StandardLogger)
}

// StandardLogger ...
type StandardLogger struct {
}

// Info ...
func (StandardLogger) Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

// Debug ...
func (StandardLogger) Debug(msg string, args ...interface{}) {
	log.Printf("[DEBUG] "+msg, args...)
}

// Error ...
func (StandardLogger) Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}
