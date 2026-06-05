package parseutils

import (
	"strconv"
	"strings"
	"time"

	"github.com/Prague-Kino/scraper/internal/errors"
)

// Converts a price string in crowns to an int
//
// Example: "190 Kč" -> 190
// Example: "100 CZK" -> 100
func CrownsToInt(s string) (int, error) {
	normalised := strings.ToLower(s)
	var cleaned string

	for _, suffix := range []string{"kč", "czk"} {
		if before, ok := strings.CutSuffix(normalised, suffix); ok {
			cleaned = strings.TrimSpace(before)
			break
		}
	}

	value, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Add time (HH:MM) to a date.
//
// Expected format: 15:04
func CombineDateTime(date time.Time, timeStr string) (time.Time, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return date, &errors.InvalidTimeFormatError{
			InvalidTime: timeStr,
		}
	}

	// Combine using the date's year/month/day and the parsed hour/minute
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		t.Hour(), t.Minute(), 0, 0,
		date.Location(),
	), nil
}
