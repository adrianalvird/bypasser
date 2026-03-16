package techniques

import (
	"strings"

	"github.com/adrianalvird/bypasser/internal/version"
	)

// Capitalization applies capitalization payloads to paths
func Capitalization(url string, payloads []string) []string {
	var capitalizedPaths []string
	for _, payload := range payloads {
		capitalizedPaths = append(capitalizedPaths, strings.ReplaceAll(url, "admin", payload))
	}
	return capitalizedPaths
}

// RandomCapitalization generates random capitalization variations of a path
func RandomCapitalization(path string) []string {
	var variations []string
	runes := []rune(path)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			uppercase := string(append(append([]rune{}, runes[:i]...), append([]rune{runes[i] - 32}, runes[i+1:]...)...))
			variations = append(variations, uppercase)
		}
	}
	return variations
}
