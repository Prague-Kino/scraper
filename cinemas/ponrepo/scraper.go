package ponrepo

import (
	"time"

	"github.com/Prague-Kino/cast"
	utils "github.com/Prague-Kino/scraper/internal/parseutils"

	"github.com/gocolly/colly/v2"
)

const (
	TimeClass      = ".event-item__date"
	DirectorClass  = ".event-item__suptitle"
	FilmTitleClass = ".event-item__title"
	FilmDetails    = ".event-item__details"
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
	time := e.ChildText(TimeClass)
	director := e.ChildText(DirectorClass)
	filmName := e.ChildText(FilmTitleClass)
	details := e.ChildText(FilmDetails)
	_ = details

	film := cast.Film{
		Title: filmName,
		Director: director,
	}

	combinedDate, err := utils.CombineDateTime(date, time)
	if err == nil {
		date = combinedDate
	}

	return cast.Screening{
		Film:  film,
		Kino:  Ponrepo,
		Date:  cast.PBTime{Time: date},
		Price: DefaultPrice,
	}
}
