package stringutils

import "strings"

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func NotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Removes extra spaces and linebreaks from a string
func Squish(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
