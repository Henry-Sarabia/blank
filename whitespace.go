package whitespace

import "regexp"

// remove returns the provided string with all of the whitespace removed.
// This includes spaces, tabs, newlines, returns, and form feeds.
func remove(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, "")
}

// isBlank returns true if the provided string is empty or only consists of whitespace.
// Returns false otherwise.
func isBlank(s string) bool {
	if remove(s) == "" {
		return true
	}

	return false
}
