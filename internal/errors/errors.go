package errors

import "fmt"

type ScrapeError struct {
	Message string
	Site    string
}

func (e *ScrapeError) Error() string {
	return fmt.Sprintf("Error encountered while scraping %s: %s", e.Site, e.Message)
}
