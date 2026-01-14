package ponrepo

import (
	"time"

	"github.com/Prague-Kino/scraper/lib"

	"github.com/gocolly/colly/v2"
)

type PonrepoScraper struct{}

func (PonrepoScraper) Kino() lib.Kino {
	return Ponrepo
}

func (PonrepoScraper) Register(c *colly.Collector, screenings *[]lib.Screening) {
	c.OnHTML("#events-list .event-group", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]lib.Screening) {
	date := parsedDate(e.Attr("id"))

	e.ForEach(".event-item", func(i int, h *colly.HTMLElement) {
		*screenings = append(*screenings, scrapeScreening(h, date))
	})
}

func scrapeScreening(e *colly.HTMLElement, date time.Time) lib.Screening {
	time := e.ChildText(".event-item__date")
	director := e.ChildText(".event-item__suptitle")
	filmName := e.ChildText(".event-item__title")
	details := e.ChildText(".event-item__details")
	_ = details

	film := lib.Film{Title: filmName, Director: director}

	return lib.Screening{
		Date:  date,
		Time:  time,
		Film:  film,
		Kino:  Ponrepo.Name,
		Price: 0,
	}
}
