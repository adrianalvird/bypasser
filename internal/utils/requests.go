package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ParseRequestFile reads a captured request file and converts it to a slice of http.Request objects.
func ParseRequestFile(filePath string) ([]*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open request file: %v", err)
	}
	defer file.Close()

	var requests []*http.Request
	var currentRequest bytes.Buffer
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line indicates the end of a request
		if line == "" && currentRequest.Len() > 0 {
			req, err := parseSingleRequest(currentRequest.String())
			if err != nil {
				return nil, fmt.Errorf("error parsing request: %v", err)
			}
			requests = append(requests, req)
			currentRequest.Reset()
			continue
		}

		// Append line to current request buffer
		currentRequest.WriteString(line + "\n")
	}

	// Handle the last request in the file
	if currentRequest.Len() > 0 {
		req, err := parseSingleRequest(currentRequest.String())
		if err != nil {
			return nil, fmt.Errorf("error parsing request: %v", err)
		}
		requests = append(requests, req)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading request file: %v", err)
	}

	return requests, nil
}

func parseSingleRequest(requestData string) (*http.Request, error) {
	lines := strings.Split(requestData, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid request data")
	}

	// Parse the request line
	requestLine := lines[0]
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}
	method, path := parts[0], parts[1]

	// Extract headers and body
	headers := http.Header{}
	var body string
	isBody := false
	host := ""

	for _, line := range lines[1:] {
		if line == "" {
			isBody = true
			continue
		}
		if isBody {
			body += line + "\n"
		} else {
			headerParts := strings.SplitN(line, ":", 2)
			if len(headerParts) == 2 {
				key := strings.TrimSpace(headerParts[0])
				value := strings.TrimSpace(headerParts[1])
				headers.Add(key, value)
				if strings.ToLower(key) == "host" {
					host = value
				}
			}
		}
	}

	// Validate the Host header
	if host == "" {
		return nil, fmt.Errorf("missing Host header in request")
	}

	// Construct the full URL
	fullURL := fmt.Sprintf("http://%s%s", host, path)

	// Create the request object
	req, err := http.NewRequest(method, fullURL, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return req, nil
}
