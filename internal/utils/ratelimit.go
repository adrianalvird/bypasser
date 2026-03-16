package utils

import (
	"math/rand"
	"sync"
	"time"
)

// RateLimiter manages the rate of requests sent
type RateLimiter struct {
	mu          sync.Mutex
	userAgents  []string
	currentRate int
	minRate     int
	maxRate     int
	ticker      *time.Ticker
	quit        chan struct{}
}

// NewRateLimiter creates a new rate limiter
// rateLimitMin and rateLimitMax define the range for requests per second
func NewRateLimiter(rateLimitMin, rateLimitMax int, userAgents []string) *RateLimiter {
	if rateLimitMin > rateLimitMax {
		panic("rateLimitMin cannot be greater than rateLimitMax")
	}

	rl := &RateLimiter{
		userAgents:  userAgents,
		minRate:     rateLimitMin,
		maxRate:     rateLimitMax,
		currentRate: rand.Intn(rateLimitMax-rateLimitMin+1) + rateLimitMin,
		quit:        make(chan struct{}),
	}

	rl.ticker = time.NewTicker(time.Second / time.Duration(rl.currentRate))
	go rl.adjustRate()
	return rl
}

// Wait blocks until the rate limiter allows the next request and returns a random User-Agent
func (rl *RateLimiter) Wait() string {
	<-rl.ticker.C
	return rl.randomUserAgent()
}

// Stop shuts down the rate limiter
func (rl *RateLimiter) Stop() {
	close(rl.quit)
	rl.ticker.Stop()
}

// AdjustRateBasedOnResponse dynamically adjusts the request rate based on server responses
func (rl *RateLimiter) AdjustRateBasedOnResponse(statusCode int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	switch {
	case statusCode == 429: // Too Many Requests
		// Reduce rate significantly
		rl.currentRate = rl.minRate
		rl.ticker.Reset(time.Second / time.Duration(rl.currentRate))
	case statusCode >= 500: // Server errors
		// Slightly decrease rate
		if rl.currentRate > rl.minRate {
			rl.currentRate -= 1
			rl.ticker.Reset(time.Second / time.Duration(rl.currentRate))
		}
	case statusCode == 200: // Success
		// Slightly increase rate, but don't exceed maxRate
		if rl.currentRate < rl.maxRate {
			rl.currentRate += 1
			rl.ticker.Reset(time.Second / time.Duration(rl.currentRate))
		}
	default:
		// No significant adjustment for other codes
	}
}

// adjustRate dynamically adjusts the request rate within the specified range
func (rl *RateLimiter) adjustRate() {
	for {
		select {
		case <-rl.quit:
			return
		case <-time.After(10 * time.Second): // Adjust rate every 10 seconds
			rl.mu.Lock()
			newRate := rand.Intn(rl.maxRate-rl.minRate+1) + rl.minRate
			if newRate != rl.currentRate {
				rl.currentRate = newRate
				rl.ticker.Reset(time.Second / time.Duration(rl.currentRate))
			}
			rl.mu.Unlock()
		}
	}
}

// randomUserAgent selects a random User-Agent from the list
func (rl *RateLimiter) randomUserAgent() string {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.userAgents[rand.Intn(len(rl.userAgents))]
}
