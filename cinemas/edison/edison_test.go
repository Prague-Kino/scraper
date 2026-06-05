package edison

import (
	"os"
	"testing"
	"time"

	"github.com/Prague-Kino/cast"
	"github.com/Prague-Kino/scraper/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --------------------------------------

var expectedScreenings = []cast.Screening{
	{
		Film:  cast.Film{Title: "Backrooms"},
		Kino:  Edison,
		Date:  cast.PBTime{Time: time.Date(2026, time.May, 30, 0, 0, 0, 0, time.Local)},
		Price: 185,
	},
	{
		Film:  cast.Film{Title: "The Blue Trail"},
		Kino:  Edison,
		Date:  cast.PBTime{Time: time.Date(2026, time.May, 31, 0, 0, 0, 0, time.Local)},
		Price: 100,
	},
}

func TestScrapeEdisonProgram(t *testing.T) {
	html, err := os.ReadFile("testdata/test.html")
	require.NoError(t, err)

	ts := testutil.ServeHTML(t, html)
	defer ts.Close()

	c := testutil.NewTestCollector()
	var screenings []cast.Screening

	scraper := EdisonScraper{}
	scraper.Register(c, &screenings)

	err = c.Visit(ts.URL)
	require.NoError(t, err)

	require.Len(t, screenings, 2)
	testutil.AssertScreenings(t, expectedScreenings, screenings)
}

func TestScraperGetKino(t *testing.T) {
	var scraper EdisonScraper
	assert.Equal(t, Edison, scraper.Kino())
}

// --------------------------------------

func TestParseEdisonDate_Success(t *testing.T) {
	input := "Monday 5.1."
	output, err := parseEdisonDate(input)

	expected, _ := time.ParseInLocation("02-01-2006", "05-01-2026", time.Local)

	require.NoError(t, err)
	assert.Equal(t, expected, output)
}

func TestParseEdisonDate_Failure(t *testing.T) {
	input := "05/01/2026"
	output, err := parseEdisonDate(input)

	expected := time.Time{}

	require.Error(t, err)
	assert.Equal(t, expected, output)
}
