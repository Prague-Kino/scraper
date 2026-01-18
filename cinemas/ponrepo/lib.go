package ponrepo

import "time"

func parseDate(dateStr string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	return parsedDate
}
