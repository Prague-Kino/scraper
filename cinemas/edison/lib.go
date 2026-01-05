package edison

import (
	"fmt"
	"strings"
	"time"
)

// Parses a date string like "Monday 5.1." into a time.Time object
func parseEdisonDate(input string) (time.Time, error) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	dayMonth := parts[len(parts)-1]

	year := time.Now().Year()
	withYear := fmt.Sprintf("%d %s", year, dayMonth)

	layout := "2006 2.1."
	return time.ParseInLocation(layout, withYear, time.Local)
}
