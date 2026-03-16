package utils

import (
	"log"
	"os"
)

var verboseLogging bool

// SetVerbose enables verbose logging
func SetVerbose(enabled bool) {
	verboseLogging = enabled
	if enabled {
		log.SetOutput(os.Stdout)
		log.Println("Verbose logging enabled")
	} else {
		log.SetOutput(nil)
	}
}

// LogVerbose logs a message only if verbose logging is enabled
func LogVerbose(message string) {
	if verboseLogging {
		log.Println(message)
	}
}
