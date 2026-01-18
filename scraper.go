package scraper

import (
	"fmt"

	e "github.com/Prague-Kino/scraper/internal/errors"
	m "github.com/Prague-Kino/scraper/models"
	"github.com/gocolly/colly/v2"
)

func ScrapeCinema(scraper m.CinemaScraper) ([]m.Screening, error) {
	var screenings []m.Screening
	var err error

	kino := scraper.Kino()

	c := colly.NewCollector(
		colly.AllowedDomains(kino.BaseDomain),
	)

	scraper.Register(c, &screenings)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", kino.ProgramURL)
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scraping %s complete!\n", kino.BaseDomain)
	})

	c.Visit(kino.ProgramURL)

	if err != nil {
		return nil, &e.ScrapeError{
			Message: err.Error(),
			Site:    kino.ProgramURL,
		}
	}

	return screenings, nil
}

func ScrapeAllCinemas() ([]m.Screening, error) {
	var screenings []m.Screening

	for _, cinema := range AllCinemas {
		s, err := ScrapeCinema(cinema)
		if err != nil {
			return nil, err
		}
		screenings = append(screenings, s...)
	}

	return screenings, nil
}
