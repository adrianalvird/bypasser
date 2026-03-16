package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// NewHTTPClient creates an HTTP client with a custom timeout and optional proxy.
func NewHTTPClient(timeout int, proxyAddress string) *http.Client {
	transport := &http.Transport{}
	if proxyAddress != "" {
		proxyURL, err := url.Parse(proxyAddress)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		} else {
			LogVerbose(fmt.Sprintf("Failed to parse proxy URL: %v", err))
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
}

// ExecuteRequest handles making an HTTP request.
func ExecuteRequest(client *http.Client, rawURL string, headers map[string]string, protocol string) (bool, string) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		LogVerbose(fmt.Sprintf("Failed to create request: %v", err))
		return false, ""
	}

	// Set headers for the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set the HTTP protocol version if specified
	if protocol != "" {
		if protocol == "HTTP/1.1" {
			req.Proto = "HTTP/1.1"
		} else if protocol == "HTTP/2" {
			req.Proto = "HTTP/2.0"
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		LogVerbose(fmt.Sprintf("Request failed: %v", err))
		return false, ""
	}
	defer resp.Body.Close()

	// Log response details if verbose logging is enabled
	body, _ := ioutil.ReadAll(resp.Body)
	LogVerbose(fmt.Sprintf("Response: %d, Body: %s", resp.StatusCode, string(body)))

	// Return success and cURL command if the response status is OK
	if resp.StatusCode == http.StatusOK {
		curlCommand := fmt.Sprintf("curl -X GET '%s'", rawURL)
		return true, curlCommand
	}
	return false, ""
}
