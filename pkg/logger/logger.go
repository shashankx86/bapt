package logger

import (
	"io"
	"log"
	"os"
)

var (
	// InfoLogger logs general info messages.
	InfoLogger *log.Logger
	// WarningLogger logs warning messages.
	WarningLogger *log.Logger
	// ErrorLogger logs error messages.
	ErrorLogger *log.Logger
)

// Initialize initializes the loggers.
func Initialize(logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	InfoLogger = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
