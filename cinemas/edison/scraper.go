package edison

import (
	"time"

	"github.com/Prague-Kino/cast"
	utils "github.com/Prague-Kino/scraper/internal/parseutils"
	su "github.com/Prague-Kino/scraper/internal/stringutils"

	"github.com/gocolly/colly/v2"
)

const (
	TimeClass     = ".time"
	FilmNameClass = ".name"
	PriceClass    = ".ticket"
)

type EdisonScraper struct{}

func (EdisonScraper) Kino() cast.Kino {
	return Edison
}

func (EdisonScraper) Register(c *colly.Collector, screenings *[]cast.Screening) {
	currentDate := time.Now()

	c.OnHTML(".program_table .line", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings, &currentDate)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]cast.Screening, currentDate *time.Time) {
	// check if line is a date header
	dateString := e.ChildText(".den")
	if su.NotEmpty(dateString) {
		processDate(dateString, currentDate)
		return
	}

	screening := parseScreening(e, *currentDate)
	*screenings = append(*screenings, screening)
}

// Parses a date string like Monday 5.1 into a time.Time object
// and updates the currentDate pointer
func processDate(dateString string, currentDate *time.Time) {
	parsedDate, err := parseEdisonDate(dateString)
	if err != nil {
		return
	}

	*currentDate = parsedDate
}

// Parses a single screening row and returns a Screening struct
func parseScreening(e *colly.HTMLElement, date time.Time) cast.Screening {
	time := e.ChildText(TimeClass)
	movieName := e.ChildText(FilmNameClass)
	priceString := e.ChildText(PriceClass)

	price, err := utils.CrownsToInt(priceString)
	if err != nil {
		price = 1
	}

	combinedDate, err := utils.CombineDateTime(date, time)
	if err == nil {
		date = combinedDate
	}

	return cast.Screening{
		Film:  cast.Film{Title: movieName},
		Kino:  Edison,
		Date:  cast.PBTime{Time: date},
		Price: price,
	}
}
