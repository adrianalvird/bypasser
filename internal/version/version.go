package version

import (
	"fmt"
	"net/http"
	"os"
	"strings"

)

const (
	CurrentVersion = "1.0.9"
	MinVersion     = "1.0.0"
	VersionURL     = "https://raw.githubusercontent.com/adrianalvird/bypasser/main/VERSION"
)

// CheckVersion verifies if the tool can run
func CheckVersion() {
	// Try to fetch the minimum version from GitHub
	resp, err := http.Get(VersionURL)
	if err != nil {
		fmt.Println("\n Cannot verify tool version - offline or repo deleted")
		fmt.Println("This tool requires version verification to run")
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("\n Version check failed - repository may be deleted")
		os.Exit(1)
	}

	// Read the minimum version from GitHub
	var remoteMinVersion string
	fmt.Fscanf(resp.Body, "%s", &remoteMinVersion)
	
	// Compare versions
	if !isVersionCompatible(CurrentVersion, remoteMinVersion) {
		fmt.Printf("\n❌ Tool version (%s) is outdated\n", CurrentVersion)
		fmt.Printf("Minimum required version: %s\n", remoteMinVersion)
		fmt.Println("Please update the tool")
		os.Exit(1)
	}
	
	// Optional: Show success message
	fmt.Printf(" Version check passed (v%s)\n", CurrentVersion)
}

func isVersionCompatible(current, min string) bool {
	// Simple version comparison (format: x.x.x)
	currentParts := strings.Split(current, ".")
	minParts := strings.Split(min, ".")
	
	for i := 0; i < len(minParts); i++ {
		if i >= len(currentParts) {
			return false
		}
		if currentParts[i] < minParts[i] {
			return false
		}
		if currentParts[i] > minParts[i] {
			return true
		}
	}
	return true
}
