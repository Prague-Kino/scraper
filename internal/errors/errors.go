package errors

import "fmt"

type ScrapeError struct {
	Message string
	Site    string
}

func (e *ScrapeError) Error() string {
	return fmt.Sprintf("Error encountered while scraping %s: %s", e.Site, e.Message)
}

type InvalidTimeFormatError struct {
	InvalidTime string
}

func (e *InvalidTimeFormatError) Error() string {
	return fmt.Sprintf(
		"Invalid time format: [%s]\n"+
		"Expected format: HH:MM (example: 15:04)",
		e.InvalidTime,
	)
}