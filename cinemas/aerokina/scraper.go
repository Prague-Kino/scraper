package aero

import (
	"time"

	"github.com/Prague-Kino/cast"
	utils "github.com/Prague-Kino/scraper/internal/parseutils"
	su "github.com/Prague-Kino/scraper/internal/stringutils"

	"github.com/gocolly/colly/v2"
)

const (
	MovieNameClass  = ".program__movie-name"
	TimeClass       = ".program__hour"
	CinemaNameClass = ".program__place--desktop"
	PriceClass      = ".program__price form"
)

type AeroScraper struct{}

func (AeroScraper) Kino() cast.Kino {
	return *Aero
}

func (AeroScraper) Register(c *colly.Collector, screenings *[]cast.Screening) {
	c.OnHTML("#program .program", func(e *colly.HTMLElement) {
		scrapeAeroProgram(e, screenings)
	})
}

func scrapeAeroProgram(e *colly.HTMLElement, screenings *[]cast.Screening) {
	var screeningDate time.Time

	// get dates
	e.ForEach("a[id^='program-day-']", func(_ int, a *colly.HTMLElement) {
		id := a.Attr("id")
		date, err := idToDate(id)
		if err == nil {
			screeningDate = date
		}
	})

	// individual screening data
	e.ForEach(".program__info-row", func(i int, row *colly.HTMLElement) {
		screening := parseScreening(row, screeningDate)
		*screenings = append(*screenings, screening)
	})
}

// Parses a single screening row and returns a Screening struct
func parseScreening(row *colly.HTMLElement, date time.Time) cast.Screening {
	movieName := row.ChildText(MovieNameClass)
	programHour := row.ChildText(TimeClass)
	cinemaName := row.ChildText(CinemaNameClass)
	priceString := row.ChildText(PriceClass)

	cinemaName = filterCinemaName(su.Squish(cinemaName))
	
	price, err := utils.CrownsToInt(priceString)
	if err != nil {
		price = 0
	}

	combinedDate, err := utils.CombineDateTime(date, programHour)
	if err == nil {
		date = combinedDate
	}

	film := cast.Film{Title: movieName}
	screening := cast.Screening{
		Film:  film,
		Kino:  getKinoByName(cinemaName),
		Date:  cast.PBTime{Time: date},
		Price: price,
	}
	return screening
}

// <div #program> contains all the screening
//
//	<div .program> program for a single day - sometimes has a single screening or multiple screenings
//	<a class="program__day" id="program-day-dd-mm-yyyy" href="#program-day-dd-mm-yyyy"> a tag with the day's date
//		<span> children with the date written as Today | Tomorrow | Sa 03/01 - the id is a better source because it spells out the whole date
//	<div.program__info> parent div for the movie screenings
//		<div.program__info-row> div for every individual screening
//			<div.program__info-row-left> contains the main screening data
//				<div.program__shared-hover> container for the movie name, screening time, and cinema
//					<div.program__info-row-left-top> container for the time and cinema
//						<div.program__hour> simply contains the time in 24h format. example: 20:00
//						<div.program__place> container for a <span> that holds the name of the cinema
//					<div.program__movie-name> this div contains the movie name inside it AND a brief description in a data tag "data-movie-desc"
//			<div.program__price> container for the price tag and buy button
//				form button.program__ticket.program__ticket--kino-aero span : the actual price of the ticket
