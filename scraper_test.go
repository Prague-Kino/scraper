// this test package actually scrapes the real websites

package scraper

import (
	"testing"

	"github.com/Prague-Kino/scraper/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type kinoScrapeTestCase struct {
	name    string
	scraper models.CinemaScraper
}

func TestScrape(t *testing.T) {
	tests := []kinoScrapeTestCase{
		{
			name:    "ScrapeCinema - Aero",
			scraper: Aero,
		},
		{
			name:    "ScrapeCinema - Edison",
			scraper: Edison,
		},
		{
			name:    "ScrapeCinema - Ponrepo",
			scraper: Ponrepo,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testScrape(tc))
	}
}

func testScrape(tc kinoScrapeTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		screenings, err := ScrapeCinema(tc.scraper)

		require.NoError(t, err)
		assert.NotEmpty(t, screenings)
	}
}

func TestScrapeAllCinemas(t *testing.T) {
	screenings, err := ScrapeAllCinemas()

	require.NoError(t, err)
	assert.NotEmpty(t, screenings)
}
