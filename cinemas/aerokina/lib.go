package aero

import (
	"strings"
	"time"
)

var cinemaNameFilters = []string{
	"Cinema",
	"Great hall",
	"Small hall",
	"Third hall",
}

// Converts a program id string in the format program-day-07-06-2026 to time.Time
func idToDate(id string) (time.Time, error) {
	datePart := strings.TrimPrefix(id, "program-day-")

	t, err := time.Parse("02-01-2006", datePart)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// Filters known substrings from cinema names so that only the main name remains
//
// Example: "Aero Cinema" -> "Aero"
//
// Example: "Lucerna Great Hall" -> "Lucerna"
func filterCinemaName(name string) string {
	s := name
	for _, filter := range cinemaNameFilters {
		s = strings.ReplaceAll(s, filter, "")
	}
	return s
}
