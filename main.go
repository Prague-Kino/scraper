package main

import (
	"fmt"

	aero "prague-kino/cinemas/aerokina"
	"prague-kino/lib"

	"github.com/gocolly/colly/v2"
)

func main() {
	var aero aero.AeroScraper

	screenings := scrapeCinema(aero)
	fmt.Println("Aerokina screenings:")
	fmt.Println(screenings)
}

func scrapeCinema(scraper lib.CinemaScraper) []lib.Screening {
	kino := scraper.Kino()

	c := colly.NewCollector(
		colly.AllowedDomains(kino.BaseDomain),
	)

	var screenings []lib.Screening

	scraper.Register(c, &screenings)

	c.Visit(kino.ProgramURL)

	return screenings
}
