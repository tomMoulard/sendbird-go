// Package strconv formats primitive types to strings.
package strconv

import "strings"

// FormatSliceToCSV converts a slice of strings to a Comma Separated Values
// string.
func FormatSliceToCSV(s []string) string {
	return strings.Join(s, ",")
}
