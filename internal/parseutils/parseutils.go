package parseutils

import (
	"strconv"
	"strings"
)

// Removes extra spaces and linebreaks from a string
func Squish(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Converts a price string in crowns to an int
//
// Example: "190 Kč" -> 190
func CrownsToInt(s string) (int, error) {
	normalised := strings.ToLower(s)
	cleaned := strings.TrimSpace(strings.TrimSuffix(normalised, "kč"))

	value, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func NotEmpty(s string) bool {
	return !IsEmpty(s)
}
