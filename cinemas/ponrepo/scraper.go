package ponrepo

import (
	"time"

	"github.com/Prague-Kino/cast/cast"

	"github.com/gocolly/colly/v2"
)

type PonrepoScraper struct{}

func (PonrepoScraper) Kino() cast.Kino {
	return Ponrepo
}

func (PonrepoScraper) Register(c *colly.Collector, screenings *[]cast.Screening) {
	c.OnHTML("#events-list .event-group", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]cast.Screening) {
	date := parseDate(e.Attr("id"))

	e.ForEach(".event-item", func(i int, h *colly.HTMLElement) {
		*screenings = append(*screenings, scrapeScreening(h, date))
	})
}

func scrapeScreening(e *colly.HTMLElement, date time.Time) cast.Screening {
	time := e.ChildText(".event-item__date")
	director := e.ChildText(".event-item__suptitle")
	filmName := e.ChildText(".event-item__title")
	details := e.ChildText(".event-item__details")
	_ = details

	film := cast.Film{Title: filmName, Director: director}

	return cast.NewScreening(
		film,
		Ponrepo.Name,
		date,
		time,
		0,
	)
}
