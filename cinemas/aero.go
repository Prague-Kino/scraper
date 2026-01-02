// Kino Aero program:
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
package cinemas

import (
	"fmt"
	"prague-kino/lib"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const KINO_NAME = "Aero"
const BASE_DOMAIN = "www.kinoaero.cz"
const PROGRAM_URL = "https://www.kinoaero.cz/en?cinema=1&sort=sort-by-data&english-friendly=1"

type AeroScraper struct{}

var Aero = lib.Kino{Name: KINO_NAME, BaseDomain: "www.kinoaero.cz", ProgramURL: PROGRAM_URL}

func (AeroScraper) Kino() lib.Kino {
	return Aero
}

func (AeroScraper) Register(c *colly.Collector, screenings *[]lib.Screening) {
	c.OnHTML("#program .program", func(e *colly.HTMLElement) {
		scrapeAeroProgram(e, screenings)
	})
}

func scrapeAeroProgram(e *colly.HTMLElement, screenings *[]lib.Screening) {
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
		movieName := row.ChildText(".program__movie-name")
		programHour := row.ChildText(".program__hour")
		price := row.ChildText(".program__price form button.program__ticket.program__ticket--kino-aero span")
		cost, _ := crownsToInt(price)
		fmt.Println(price)

		projectionId := row.ChildAttr(".program__info-row-left-top", "data-projection")
		fmt.Println("projection id: ", projectionId)

		var film lib.Film
		film = lib.Film{
			Title: movieName,
		}

		screening := lib.Screening{
			Film: film,
			Kino: KINO_NAME,
			Date: screeningDate,
			Cost: cost,
			Time: programHour,
		}
		fmt.Println(screening)
		*screenings = append(*screenings, screening)
	})
}

// Converts a program id string in the format program-day-07-06-2026 to time.Time
func idToDate(id string) (time.Time, error) {
	datePart := strings.TrimPrefix(id, "program-day-")

	t, err := time.Parse("02-01-2006", datePart)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// Converts a price string in crowns to an int
//
// Example: "190 Kč" -> 190
func crownsToInt(s string) (int, error) {
	cleaned := strings.TrimSpace(strings.TrimSuffix(s, "Kč"))

	value, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0, err
	}
	return value, nil
}
