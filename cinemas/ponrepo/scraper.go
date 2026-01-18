package ponrepo

import (
	"time"

	"github.com/Prague-Kino/scraper/models"

	"github.com/gocolly/colly/v2"
)

type PonrepoScraper struct{}

func (PonrepoScraper) Kino() models.Kino {
	return Ponrepo
}

func (PonrepoScraper) Register(c *colly.Collector, screenings *[]models.Screening) {
	c.OnHTML("#events-list .event-group", func(e *colly.HTMLElement) {
		scrapeProgram(e, screenings)
	})
}

func scrapeProgram(e *colly.HTMLElement, screenings *[]models.Screening) {
	date := parseDate(e.Attr("id"))

	e.ForEach(".event-item", func(i int, h *colly.HTMLElement) {
		*screenings = append(*screenings, scrapeScreening(h, date))
	})
}

func scrapeScreening(e *colly.HTMLElement, date time.Time) models.Screening {
	time := e.ChildText(".event-item__date")
	director := e.ChildText(".event-item__suptitle")
	filmName := e.ChildText(".event-item__title")
	details := e.ChildText(".event-item__details")
	_ = details

	film := models.Film{Title: filmName, Director: director}

	return models.Screening{
		Date:  date,
		Time:  time,
		Film:  film,
		Kino:  Ponrepo.Name,
		Price: 0,
	}
}
