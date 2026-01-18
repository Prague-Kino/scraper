package edison

import (
	"time"

	utils "github.com/Prague-Kino/scraper/internal/parseutils"
	"github.com/Prague-Kino/scraper/models"

	"github.com/gocolly/colly/v2"
)

type EdisonScraper struct{}

func (EdisonScraper) Kino() models.Kino {
	return Edison
}

func (EdisonScraper) Register(c *colly.Collector, screenings *[]models.Screening) {
	currentDate := time.Now()

	c.OnHTML(".program_table .line", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings, &currentDate)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]models.Screening, currentDate *time.Time) {
	// check if line is a date header
	dateString := e.ChildText(".den")
	if utils.NotEmpty(dateString) {
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
func parseScreening(e *colly.HTMLElement, date time.Time) models.Screening {
	time := e.ChildText(".time")
	movieName := e.ChildText(".name")
	priceString := e.ChildText(".ticket")

	price, err := utils.CrownsToInt(priceString)
	if err != nil {
		price = 0
	}

	return models.NewScreening(
		models.Film{Title: movieName},
		Edison.Name,
		date,
		time,
		price,
	)
}
