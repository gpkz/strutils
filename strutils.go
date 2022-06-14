//Package strutils provides some simple string modifications
package strutils

import "strings"

// StandardizeSpaces replaces multiple spaces with a single space, also removing returns and tabs
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// RemoveNonAlphaChars removes all non-alpha characters from the string (except for spaces)
func RemoveNonAlphaChars(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
