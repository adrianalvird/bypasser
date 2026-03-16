package bypass

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianalvird/bypasser/payloads"
)

// BypassResult holds information about a successful bypass
type BypassResult struct {
	Technique   string
	CurlCommand string
}

// GetUserAgents returns the list of user agents from payloads
func GetUserAgents() []string {
	return payloads.Payloads.UserAgents
}

// Bypass applies multiple bypass techniques using embedded payloads
func Bypass(client *http.Client, rawURL string, runContinuously bool) []BypassResult {
	var results []BypassResult
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("Invalid URL: %s, error: %v\n", rawURL, err)
		return results
	}

	baseURL := parsedURL.Host
	basePath := parsedURL.Path

	// Helper function to add result and check if we should continue
	addResult := func(technique, curlCommand string) bool {
		results = append(results, BypassResult{
			Technique:   technique,
			CurlCommand: curlCommand,
		})
		return runContinuously // Continue if in continuous mode
	}

	// Path Manipulation
	for _, path := range payloads.Payloads.Paths {
		modifiedPath := strings.Replace(path, "{{baseURL}}", basePath, -1)
		modifiedURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, baseURL, modifiedPath)
		if success, curlCommand := sendRequest(client, modifiedURL, "GET", nil, nil); success {
			if !addResult(fmt.Sprintf("Path Manipulation: %s", path), curlCommand) {
				return results
			}
		}
	}

	// Header Modification
	for _, header := range payloads.Payloads.Headers {
		modifiedHeader := replacePlaceholders(header, baseURL)
		if success, curlCommand := sendRequest(client, rawURL, "GET", modifiedHeader, nil); success {
			if !addResult(fmt.Sprintf("Header Modification: %+v", header), curlCommand) {
				return results
			}
		}
	}

	// Header Combinations
	for _, combination := range payloads.Payloads.HeaderCombinations {
		modifiedHeaders := replacePlaceholders(combination["headers"], baseURL)
		if success, curlCommand := sendRequest(client, rawURL, "GET", modifiedHeaders, nil); success {
			if !addResult(fmt.Sprintf("Header Combination Attack: %+v", modifiedHeaders), curlCommand) {
				return results
			}
		}
	}

	// Extension Injection
	for _, ext := range payloads.Payloads.ExtensionInjection {
		modifiedURL := fmt.Sprintf("%s://%s%s%s", parsedURL.Scheme, baseURL, basePath, ext)
		if success, curlCommand := sendRequest(client, modifiedURL, "GET", nil, nil); success {
			if !addResult(fmt.Sprintf("Extension Injection: %s", ext), curlCommand) {
				return results
			}
		}
	}

	// Protocol Manipulation
	for _, protocol := range payloads.Payloads.Protocols {
		if success, curlCommand := sendRequestWithProtocol(client, rawURL, protocol); success {
			if !addResult(fmt.Sprintf("Protocol Change: %s", protocol), curlCommand) {
				return results
			}
		}
	}

	// Content-Length Variation
	for _, cl := range payloads.Payloads.ContentLength {
		headers := replacePlaceholders(cl, baseURL)
		if success, curlCommand := sendRequest(client, rawURL, "POST", headers, map[string]string{"data": "test"}); success {
			if !addResult(fmt.Sprintf("Content-Length Variation: %+v", headers), curlCommand) {
				return results
			}
		}
	}

	// Path + Header Combination
	for _, path := range payloads.Payloads.Paths {
		modifiedPath := strings.Replace(path, "{{baseURL}}", basePath, -1)
		modifiedURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, baseURL, modifiedPath)

		for _, header := range payloads.Payloads.Headers {
			modifiedHeader := replacePlaceholders(header, baseURL)
			if success, curlCommand := sendRequest(client, modifiedURL, "GET", modifiedHeader, nil); success {
				if !addResult(fmt.Sprintf("Path + Header Combination: Path: %s, Header: %+v", path, header), curlCommand) {
					return results
				}
			}
		}

		for _, headerCombo := range payloads.Payloads.HeaderCombinations {
			modifiedHeaders := replacePlaceholders(headerCombo["headers"], baseURL)
			if success, curlCommand := sendRequest(client, modifiedURL, "GET", modifiedHeaders, nil); success {
				if !addResult(fmt.Sprintf("Path + Multi-Header Combination: Path: %s, Headers: %+v", path, modifiedHeaders), curlCommand) {
					return results
				}
			}
		}
	}

	// Header + Extension Injection
	for _, ext := range payloads.Payloads.ExtensionInjection {
		modifiedURL := fmt.Sprintf("%s://%s%s%s", parsedURL.Scheme, baseURL, basePath, ext)
		for _, header := range payloads.Payloads.Headers {
			modifiedHeader := replacePlaceholders(header, baseURL)
			if success, curlCommand := sendRequest(client, modifiedURL, "GET", modifiedHeader, nil); success {
				if !addResult(fmt.Sprintf("Header + Extension Injection: URL: %s, Header: %+v", modifiedURL, header), curlCommand) {
					return results
				}
			}
		}
	}

	// Header + Protocol Downgrade
	for _, protocol := range payloads.Payloads.Protocols {
		req, err := http.NewRequest("GET", rawURL, nil)
		if err != nil {
			continue
		}
		req.Proto = protocol
		req.ProtoMajor, req.ProtoMinor = parseProtocol(protocol)

		for _, header := range payloads.Payloads.Headers {
			for key, value := range header {
				req.Header.Set(key, value)
			}
		}

		if success, curlCommand := executeRequest(client, req); success {
			if !addResult(fmt.Sprintf("Header + Protocol Downgrade: Protocol: %s, Headers: %+v", protocol, req.Header), curlCommand) {
				return results
			}
		}
	}

	// Header + Protocol + Extension Injection
	for _, ext := range payloads.Payloads.ExtensionInjection {
		for _, protocol := range payloads.Payloads.Protocols {
			modifiedURL := fmt.Sprintf("%s://%s%s%s", protocol, baseURL, basePath, ext)
			for _, header := range payloads.Payloads.Headers {
				modifiedHeader := replacePlaceholders(header, baseURL)
				if success, curlCommand := sendRequest(client, modifiedURL, "GET", modifiedHeader, nil); success {
					if !addResult(fmt.Sprintf("Header + Protocol Downgrade + Extension Injection: URL: %s, Protocol: %s, Header: %+v", modifiedURL, protocol, modifiedHeader), curlCommand) {
						return results
					}
				}
			}
		}
	}

	// Multi Header + Path + Protocol
	for _, path := range payloads.Payloads.Paths {
		modifiedPath := strings.Replace(path, "{{baseURL}}", basePath, -1)
		for _, protocol := range payloads.Payloads.Protocols {
			modifiedURL := fmt.Sprintf("%s://%s%s", protocol, baseURL, modifiedPath)
			for _, headerCombo := range payloads.Payloads.HeaderCombinations {
				modifiedHeaders := replacePlaceholders(headerCombo["headers"], baseURL)
				if success, curlCommand := sendRequest(client, modifiedURL, "GET", modifiedHeaders, nil); success {
					if !addResult(fmt.Sprintf("Multi Header + Path Manipulation + Protocol Downgrade: URL: %s, Protocol: %s, Headers: %+v", modifiedURL, protocol, modifiedHeaders), curlCommand) {
						return results
					}
				}
			}
		}
	}

	return results
}

// replacePlaceholders replaces placeholders in the payloads with actual values
func replacePlaceholders(payload map[string]string, baseURL string) map[string]string {
	modified := make(map[string]string)
	for key, value := range payload {
		modified[key] = strings.Replace(value, "{{baseURL}}", baseURL, -1)
	}
	return modified
}

// BypassWithRequest applies techniques to a captured HTTP request
func BypassWithRequest(client *http.Client, req *http.Request, runContinuously bool) []BypassResult {
	var results []BypassResult

	addResult := func(technique, curlCommand string) bool {
		results = append(results, BypassResult{
			Technique:   technique,
			CurlCommand: curlCommand,
		})
		return runContinuously
	}

	for _, header := range payloads.Payloads.Headers {
		origHeaders := req.Header.Clone()
		for key, value := range header {
			req.Header.Set(key, value)
		}
		success, curlCommand := executeRequest(client, req)
		req.Header = origHeaders // Restore original headers
		
		if success {
			if !addResult(fmt.Sprintf("Header Modification: %+v", header), curlCommand) {
				return results
			}
		}
	}

	for _, method := range payloads.Payloads.Methods {
		origMethod := req.Method
		req.Method = method
		success, curlCommand := executeRequest(client, req)
		req.Method = origMethod // Restore original method
		
		if success {
			if !addResult(fmt.Sprintf("Method Change: %s", method), curlCommand) {
				return results
			}
		}
	}

	return results
}

// sendRequest sends an HTTP request and returns the success status and curl command for reproduction
func sendRequest(client *http.Client, rawURL, method string, headers map[string]string, body map[string]string) (bool, string) {
	req, err := http.NewRequest(method, rawURL, nil)
	if err != nil {
		return false, ""
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		curlCommand := generateCurlCommand(method, rawURL, headers, body)
		return true, curlCommand
	}

	return false, ""
}

// sendRequestWithProtocol sends a request with a custom protocol (HTTP/1.1, HTTP/2, etc.)
func sendRequestWithProtocol(client *http.Client, rawURL, protocol string) (bool, string) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return false, ""
	}

	req.Proto = protocol
	req.ProtoMajor, req.ProtoMinor = parseProtocol(protocol)

	resp, err := client.Do(req)
	if err != nil {
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		curlCommand := generateCurlFromRequest(req)
		return true, curlCommand
	}

	return false, ""
}

// parseProtocol parses a protocol string into its major and minor version numbers
func parseProtocol(protocol string) (int, int) {
	switch protocol {
	case "HTTP/0.9":
		return 0, 9
	case "HTTP/1.0":
		return 1, 0
	case "HTTP/1.1":
		return 1, 1
	case "HTTP/2.0", "HTTP/2":
		return 2, 0
	case "HTTP/3.0", "HTTP/3":
		return 3, 0
	default:
		return 1, 1
	}
}

// executeRequest sends an HTTP request and evaluates the response
func executeRequest(client *http.Client, req *http.Request) (bool, string) {
	resp, err := client.Do(req)
	if err != nil {
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		curlCommand := generateCurlFromRequest(req)
		return true, curlCommand
	}

	return false, ""
}

// generateCurlCommand generates a curl command for reproduction
func generateCurlCommand(method, url string, headers map[string]string, data map[string]string) string {
	curl := fmt.Sprintf("curl -X %s \"%s\"", method, url)

	for key, value := range headers {
		curl += fmt.Sprintf(" -H \"%s: %s\"", key, value)
	}

	if len(data) > 0 {
		curl += " -d '"
		for key, value := range data {
			curl += fmt.Sprintf("%s=%s&", key, value)
		}
		curl = strings.TrimRight(curl, "&") + "'"
	}

	return curl
}

// generateCurlFromRequest generates a curl command from an HTTP request object
func generateCurlFromRequest(req *http.Request) string {
	curl := fmt.Sprintf("curl -X %s \"%s\"", req.Method, req.URL.String())
	for key, values := range req.Header {
		for _, value := range values {
			curl += fmt.Sprintf(" -H \"%s: %s\"", key, value)
		}
	}
	return curl
}
