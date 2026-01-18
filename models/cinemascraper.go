package models

import (
	"github.com/Prague-Kino/cast/cast"
	"github.com/gocolly/colly/v2"
)

type CinemaScraper interface {
	Kino() cast.Kino
	Register(c *colly.Collector, screenings *[]cast.Screening)
}
