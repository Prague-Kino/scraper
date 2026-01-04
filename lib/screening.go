package lib

import (
	"fmt"
	"time"
)

type Screening struct {
	Film Film
	Kino string
	Date time.Time
	Time string
	Cost int
}

func (s Screening) String() string {
	return fmt.Sprintf(
		"{ %-10s %5s | %-25s @ %-6s for %4d Kč }\n",
		s.Kino,
		s.Date.Format("01/02"),
		s.Film.Title,
		s.Time,
		s.Cost,
	)
}

func NewScreening(film Film, kino string, date time.Time, timeStr string, cost int) Screening {
	return Screening{
		Film: film,
		Kino: kino,
		Date: date,
		Time: timeStr,
		Cost: cost,
	}
}
