package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// InitLogger initializes the loggers
func InitLogger() {
	infoFile, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open info log file:", err)
	}

	errorFile, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open error log file:", err)
	}

	debugFile, err := os.OpenFile("logs/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open debug log file:", err)
	}

	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(debugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs an info message
func LogInfo(message string) {
	if InfoLogger != nil {
		InfoLogger.Println(message)
	}
}

// LogError logs an error message
func LogError(message string) {
	if ErrorLogger != nil {
		ErrorLogger.Println(message)
	}
}

// LogDebug logs a debug message
func LogDebug(message string) {
	if DebugLogger != nil {
		DebugLogger.Println(message)
	}
}

// LogRequest logs an HTTP request
func LogRequest(method, path string, duration time.Duration, statusCode int) {
	message := fmt.Sprintf("%s %s - %d - %v", method, path, statusCode, duration)
	if statusCode >= 400 {
		LogError(message)
	} else {
		LogInfo(message)
	}
}
