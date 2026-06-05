package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Prague-Kino/cast"
	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
)

// Create a collector that's allowed to hit localhost test servers
func NewTestCollector() *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains("127.0.0.1"),
	)
}

// Spin up a test server with given HTML, return its URL
func ServeHTML(t *testing.T, html []byte) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(html)
	}))
}

func AssertDate(t *testing.T, actual, expected cast.PBTime) {
	t.Helper()

	assert.Equal(
		t,
		expected.Time.Format("2006-01-02 15:04"),
		actual.Time.Format("2006-01-02 15:04"),
	)
}

func AssertScreenings(
	t *testing.T,
	expectedScreenings, screenings []cast.Screening,
) {
	t.Helper()
	
	for i, s := range screenings {
		assert.Equal(t, expectedScreenings[i].Film.Title, s.Film.Title)
		assert.Equal(t, expectedScreenings[i].Kino.Name, s.Kino.Name)
		assert.Equal(t, expectedScreenings[i].Price, s.Price)
		assert.Equal(t, expectedScreenings[i].Date, s.Date)
	}
}
