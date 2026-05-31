package aero

import (
	"os"
	"testing"
	"time"

	"github.com/Prague-Kino/cast"
	"github.com/Prague-Kino/scraper/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --------------------------

var expectedScreenings = []cast.Screening{
	{
		Film:  cast.Film{Title: "The Devil Wears Prada 2"},
		Kino:  cast.Kino{Name: "Kino Aero"},
		Date:  cast.PBTime{Time: time.Date(2026, time.June, 4, 13, 30, 0, 0, time.UTC)},
		Price: 90,
	},
	{
		Film:  cast.Film{Title: "Backrooms"},
		Kino:  cast.Kino{Name: "Kino Aero"},
		Date:  cast.PBTime{Time: time.Date(2026, time.June, 4, 18, 0, 0, 0, time.UTC)},
		Price: 180,
	},
}

func TestScrapeAeroProgram(t *testing.T) {
	html, err := os.ReadFile("testdata/test.html")
	require.NoError(t, err)

	ts := testutil.ServeHTML(t, html)
	defer ts.Close()

	c := testutil.NewTestCollector()
	var screenings []cast.Screening

	scraper := AeroScraper{}
	scraper.Register(c, &screenings)

	err = c.Visit(ts.URL)
	require.NoError(t, err)

	require.Len(t, screenings, 2)

	for i, s := range screenings {
		assert.Equal(t, expectedScreenings[i].Film.Title, s.Film.Title)
		assert.Equal(t, expectedScreenings[i].Kino.Name, s.Kino.Name)
		assert.Equal(t, expectedScreenings[i].Price, s.Price)
		assert.Equal(t, expectedScreenings[i].Date, s.Date)
	}
}

func TestScraperGetKino(t *testing.T) {
	var scraper AeroScraper
	assert.Equal(t, *Aero, scraper.Kino())
}

// ---------------------------------

func TestGetKinoByName_Valid(t *testing.T) {
	assert.Equal(t, *Aero, getKinoByName("Aero"))
	assert.Equal(t, *Svetozor, getKinoByName("Světozor"))
	assert.Equal(t, *Lucerna, getKinoByName("Lucerna"))
	assert.Equal(t, *BigOko, getKinoByName("Bio Oko"))
	assert.Equal(t, *Pritomnost, getKinoByName("Přítomnost"))
}

func TestGetKinoByName_Invalid(t *testing.T) {
	assert.Equal(t, *Aero, getKinoByName("invalid name"))
	assert.Equal(t, *Aero, getKinoByName("inbeignq93gn-"))
	assert.Equal(t, *Aero, getKinoByName(""))
}

// ---------------------------------

func TestIdToDate_Success(t *testing.T) {
	input := "program-day-07-06-2026"
	output, err := idToDate(input)

	expected, _ := time.Parse("02-01-2006", "07-06-2026")

	require.NoError(t, err)
	assert.Equal(t, expected, output)
}

func TestIdToDate_Failure(t *testing.T) {
	input := "program-day-7/6/26"
	output, err := idToDate(input)

	expected := time.Time{}

	require.Error(t, err)
	assert.Equal(t, expected, output)
}