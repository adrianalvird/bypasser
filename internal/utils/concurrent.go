package utils

import (
	"fmt"
	"sync"
)

// ConcurrentManager manages concurrent requests
type ConcurrentManager struct {
	rateLimiter *RateLimiter
	numWorkers  int
}

// NewConcurrentManager initializes a new ConcurrentManager
func NewConcurrentManager(rateLimiter *RateLimiter, numWorkers int) *ConcurrentManager {
	return &ConcurrentManager{
		rateLimiter: rateLimiter,
		numWorkers:  numWorkers,
	}
}

// ExecuteConcurrentRequests processes URLs with concurrent workers
func (cm *ConcurrentManager) ExecuteConcurrentRequests(urls []string, requestFunc func(string, string)) {
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup

	// Fill the job queue with URLs
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// Start workers
	for i := 0; i < cm.numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for url := range jobs {
				userAgent := cm.rateLimiter.Wait() // Respect rate limiting
				requestFunc(url, userAgent)
				fmt.Printf("[Worker %d] Processed: %s\n", workerID, url)
			}
		}(i)
	}

	wg.Wait() // Wait for all workers to finish
}
