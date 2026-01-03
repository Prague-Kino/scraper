package aero

import (
	"strconv"
	"strings"
	"time"
)

// Converts a program id string in the format program-day-07-06-2026 to time.Time
func idToDate(id string) (time.Time, error) {
	datePart := strings.TrimPrefix(id, "program-day-")

	t, err := time.Parse("02-01-2006", datePart)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// Converts a price string in crowns to an int
//
// Example: "190 Kč" -> 190
func crownsToInt(s string) (int, error) {
	cleaned := strings.TrimSpace(strings.TrimSuffix(s, "Kč"))

	value, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0, err
	}
	return value, nil
}
