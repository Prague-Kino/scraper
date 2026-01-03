package lib

import (
	"fmt"
	"time"
)

type Kino struct {
	Name       string
	BaseDomain string
	ProgramURL string
}

type Film struct {
	Title   string
	Year    int
	Runtime int
	Country string
}

type Screening struct {
	Film Film
	Kino string
	Date time.Time
	Time string
	Cost int
}

func (s Screening) String() string {
	return fmt.Sprintf(
		"{ %s \t| %s @ %s for %d Kč | %s }\n",
		s.Date.Format("01/02"),
		s.Film.Title,
		s.Time,
		s.Cost,
		s.Kino,
	)
}
