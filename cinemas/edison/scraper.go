package edison

import (
	"time"

	"github.com/Prague-Kino/scraper/lib"

	"github.com/gocolly/colly/v2"
)

type EdisonScraper struct{}

func (EdisonScraper) Kino() lib.Kino {
	return Edison
}

func (EdisonScraper) Register(c *colly.Collector, screenings *[]lib.Screening) {
	currentDate := time.Now()

	c.OnHTML(".program_table .line", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings, &currentDate)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]lib.Screening, currentDate *time.Time) {
	// check if line is a date header
	dateString := e.ChildText(".den")
	if !lib.IsEmpty(dateString) {
		processDate(dateString, currentDate)
		return
	}

	screening := parseScreening(e, *currentDate)
	*screenings = append(*screenings, screening)
}

func processDate(dateString string, currentDate *time.Time) {
	parsedDate, err := parseEdisonDate(dateString)
	if err != nil {
		return
	}

	*currentDate = parsedDate
}

func parseScreening(e *colly.HTMLElement, date time.Time) lib.Screening {
	time := e.ChildText(".time")
	movieName := e.ChildText(".name")
	priceString := e.ChildText(".ticket")

	price, err := lib.CrownsToInt(priceString)
	if err != nil {
		price = 0
	}

	return lib.NewScreening(
		lib.Film{Title: movieName},
		Edison.Name,
		date,
		time,
		price,
	)
}
