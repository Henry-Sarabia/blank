package blank

import (
	"unicode"
)

// RemoveSpace returns the provided string with all of the whitespace removed.
// This includes spaces, tabs, newlines, returns, form feeds and other
// space-like characters.
//
// For more information on what is considered whitespace, visit:
// https://golang.org/pkg/unicode/#IsSpace
func RemoveSpace(str string) string {
	var out []rune
	for _, r := range str {
		if !unicode.IsSpace(r) {
			out = append(out, r)
		}
	}

	return string(out)
}

// Is returns true if the provided string is empty or consists only of
// whitespace. Returns false otherwise.
func Is(str string) bool {
	if RemoveSpace(str) == "" {
		return true
	}

	return false
}

// Has returns true if any of the strings in the provided slice is empty
// or consists of only whitespace. If the slice is nil, returns true as well.
// Returns false otherwise.
func Has(slice []string) bool {
	if len(slice) <= 0 {
		return true
	}

	for _, s := range slice {
		if Is(s) {
			return true
		}
	}

	return false
}
