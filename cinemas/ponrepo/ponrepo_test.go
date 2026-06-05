package ponrepo

import (
	"os"
	"testing"
	"time"

	"github.com/Prague-Kino/cast"
	"github.com/Prague-Kino/scraper/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------

var expectedScreenings = []cast.Screening{
	{
		Film:  cast.Film{Title: "Gripsholm Castle / Schloss Gripsholm"},
		Kino:  Ponrepo,
		Date:  cast.PBTime{Time: time.Date(2026, time.May, 30, 18, 0, 0, 0, time.UTC)},
		Price: DefaultPrice,
	},
	{
		Film:  cast.Film{Title: "Joke"},
		Kino:  Ponrepo,
		Date:  cast.PBTime{Time: time.Date(2026, time.May, 30, 20, 30, 0, 0, time.UTC)},
		Price: DefaultPrice,
	},
}

func TestPonrepoScrapeProgram(t *testing.T) {
	html, err := os.ReadFile("testdata/test.html")
	require.NoError(t, err)

	ts := testutil.ServeHTML(t, html)
	defer ts.Close()

	c := testutil.NewTestCollector()
	var screenings []cast.Screening

	scraper := PonrepoScraper{}
	scraper.Register(c, &screenings)

	err = c.Visit(ts.URL)
	require.NoError(t, err)

	require.Len(t, screenings, 2)
	testutil.AssertScreenings(t, expectedScreenings, screenings)
}

func TestScraperGetKino(t *testing.T) {
	var scraper PonrepoScraper
	assert.Equal(t, Ponrepo, scraper.Kino())
}