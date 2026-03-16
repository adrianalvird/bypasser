package techniques

import (
	"fmt"
	"net/url"

	"github.com/adrianalvird/bypasser/internal/version"
)

// PathManipulation modifies the URL path for bypass attempts.
func PathManipulation(rawURL string, headers map[string]string, protocol string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	payloads := []string{"%2e", "%2f", "..%2f"}
	for _, payload := range payloads {
		modifiedPath := fmt.Sprintf("%s%s", payload, parsedURL.Path)
		parsedURL.Path = modifiedPath
		return parsedURL.String(), nil
	}
	return "", fmt.Errorf("no payloads applied")
}
