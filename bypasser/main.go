package main

import (
    "bufio"
    "crypto/tls"
    "flag"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"

    "github.com/adrianalvird/bypasser/internal/bypass"
    "github.com/adrianalvird/bypasser/internal/utils"
    "github.com/adrianalvird/bypasser/internal/version"
)

var (
	timeout      = flag.Int("timeout", 10, "Timeout for HTTP requests in seconds")
	proxy        = flag.String("proxy", "", "Proxy server (e.g., http://127.0.0.1:8080)")
	requestFile  = flag.String("request", "", "Path to a captured request file")
	verbose      = flag.Bool("verbose", false, "Enable verbose logging")
	static       = flag.String("static", "", "Static string to append to paths")
	rateLimit    = flag.Int("ratelimit", 10, "Maximum requests per second")
	runContinuously = flag.Bool("r", false, "Continue testing even after successful bypass")
)

func main() {
	// ONLY THIS LINE ADDED - version check first
	version.CheckVersion()
	
	printLogo()
	flag.Parse()

	// Set verbose logging
	utils.SetVerbose(*verbose)

	httpClient := configureHTTPClient(*proxy, *timeout)

	if *requestFile != "" {
		processRequestFile(httpClient, *requestFile, *runContinuously)
	} else {
		processStdin(httpClient, *runContinuously)
	}
}

func configureHTTPClient(proxy string, timeout int) *http.Client {
	var httpClient *http.Client

	if proxy != "" {
		if !strings.HasPrefix(proxy, "http://") && !strings.HasPrefix(proxy, "https://") {
			proxy = "http://" + proxy
		}

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing proxy URL: %v\n", err)
			os.Exit(1)
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyURL(proxyURL),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Duration(timeout) * time.Second,
		}
		utils.LogVerbose(fmt.Sprintf("Proxy enabled: %s", proxy))
	} else {
		httpClient = &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}
	}

	return httpClient
}

func processRequestFile(client *http.Client, filePath string, runContinuously bool) {
	requests, err := utils.ParseRequestFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing request file: %v\n", err)
		os.Exit(1)
	}

	// Get user agents from payloads
	userAgents := bypass.GetUserAgents()
	
	rateLimiter := utils.NewRateLimiter(1, *rateLimit, userAgents)
	concurrentManager := utils.NewConcurrentManager(rateLimiter, 5)

	urls := []string{}
	for _, req := range requests {
		urls = append(urls, req.URL.String())
	}

	concurrentManager.ExecuteConcurrentRequests(urls, func(rawURL, userAgent string) {
		results := bypass.Bypass(client, rawURL, runContinuously)
		if len(results) > 0 {
			for i, result := range results {
				fmt.Printf("\n[+] Bypass %d successful for: %s\n", i+1, rawURL)
				fmt.Printf("   Technique: %s\n", result.Technique)
				fmt.Printf("   Curl Command: %s\n", result.CurlCommand)
				
				// Stop after first success if not in continuous mode
				if !runContinuously {
					break
				}
			}
		} else {
			utils.LogVerbose(fmt.Sprintf("Bypass failed for: %s", rawURL))
		}
	})
}

func processStdin(client *http.Client, runContinuously bool) {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}

	for scanner.Scan() {
		rawURL := scanner.Text()
		urls = append(urls, rawURL)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
		return
	}

	// Get user agents from payloads
	userAgents := bypass.GetUserAgents()
	
	rateLimiter := utils.NewRateLimiter(1, *rateLimit, userAgents)
	concurrentManager := utils.NewConcurrentManager(rateLimiter, 5)

	concurrentManager.ExecuteConcurrentRequests(urls, func(rawURL, userAgent string) {
		results := bypass.Bypass(client, rawURL, runContinuously)
		if len(results) > 0 {
			for i, result := range results {
				fmt.Printf("\n[+] Bypass %d successful for: %s\n", i+1, rawURL)
				fmt.Printf("   Technique: %s\n", result.Technique)
				fmt.Printf("   Curl Command: %s\n", result.CurlCommand)
				
				// Stop after first success if not in continuous mode
				if !runContinuously {
					break
				}
			}
		} else {
			utils.LogVerbose(fmt.Sprintf("Bypass failed for: %s", rawURL))
		}
	})
}

func printLogo() {
	logo := `
 ______  __   __  _____  _______ _______ _______ _______  ______
 |_____]   \_/   |_____] |_____| |______ |______ |______ |_____/
 |_____]    |    |       |     | ______| ______| |______ |    \_
    by @adrianalvird			Version %s
                                                                
`
	fmt.Printf(logo, version.CurrentVersion)
}