package lib

import "github.com/gocolly/colly/v2"

type CinemaScraper interface {
	Kino() Kino
	Register(c *colly.Collector, screenings *[]Screening)
}
