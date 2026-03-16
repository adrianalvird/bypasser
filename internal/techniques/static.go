package techniques

import (
	"fmt"
	"github.com/adrianalvird/bypasser/internal/version"
	)

// StaticAppend adds a static string to a given URL path
func StaticAppend(path string, staticString string) string {
	return fmt.Sprintf("%s/%s", path, staticString)
}
