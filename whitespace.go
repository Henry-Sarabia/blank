package whitespace

import "regexp"

// Remove returns the provided string with all of the whitespace removed.
// This includes spaces, tabs, newlines, returns, and form feeds.
func Remove(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, "")
}

// IsBlank returns true if the provided string is empty or only consists of whitespace.
// Returns false otherwise.
func IsBlank(s string) bool {
	if Remove(s) == "" {
		return true
	}

	return false
}
