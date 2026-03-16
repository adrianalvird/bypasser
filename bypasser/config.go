package main

import (
	"flag"
	)

// Config holds application configurations
type Config struct {
	Timeout       int
	Verbose       bool
	Static        string
	RateLimit     int
	RunContinuously bool  // Add the new flag here
}

func ParseConfig() Config {
	// Add the new flag
	runContinuously := flag.Bool("r", false, "Continue testing even after successful bypass")
	flag.Parse()

	return Config{
		Timeout:       *timeout,
		Verbose:       *verbose,
		Static:        *static,
		RateLimit:     *rateLimit,
		RunContinuously: *runContinuously,  // Add to config
	}
}
