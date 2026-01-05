package main

import (
	"fmt"

	aero "prague-kino/cinemas/aerokina"
	ed "prague-kino/cinemas/edison"
	"prague-kino/lib"

	"github.com/gocolly/colly/v2"
)

func main() {
	var aero aero.AeroScraper
	screenings := scrapeCinema(aero)
	fmt.Println("Aerokina screenings:")
	fmt.Println(screenings)

	var edison ed.EdisonScraper
	fmt.Println("Edison screenings:")
	screenings2 := scrapeCinema(edison)
	fmt.Println(screenings2)
}

func scrapeCinema(scraper lib.CinemaScraper) []lib.Screening {
	kino := scraper.Kino()

	c := colly.NewCollector(
		colly.AllowedDomains(kino.BaseDomain),
	)

	var screenings []lib.Screening

	scraper.Register(c, &screenings)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", kino.ProgramURL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraping complete!")
	})

	c.Visit(kino.ProgramURL)

	return screenings
}
