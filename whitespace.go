package whitespace

import "regexp"

// Remove returns the provided string with all of the whitespace removed.
// This includes spaces, tabs, newlines, returns, and form feeds.
func Remove(str string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, "")
}

// IsBlank returns true if the provided string is empty or only consists of whitespace.
// Returns false otherwise.
func IsBlank(str string) bool {
	if Remove(str) == "" {
		return true
	}

	return false
}

// HasBlank returns true if any of the strings in the provided slice is empty
// or consists of whitespace. If the slice is nil, returns true as well.
// Returns false otherwise.
func HasBlank(slice []string) bool {
	if len(slice) <= 0 {
		return true
	}

	for _, s := range slice {
		if IsBlank(s) {
			return true
		}
	}
	return false
}
