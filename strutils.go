// Package strutils provides some simple string modifications
package strutils

import "strings"

// StandardizeSpaces replaces multiple spaces with a single space, also removing returns and tabs
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
