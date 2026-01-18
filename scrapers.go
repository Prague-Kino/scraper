package scraper

import (
	aero "github.com/Prague-Kino/scraper/cinemas/aerokina"
	ed "github.com/Prague-Kino/scraper/cinemas/edison"
	pon "github.com/Prague-Kino/scraper/cinemas/ponrepo"
	m "github.com/Prague-Kino/scraper/models"
)

var (
	Aero    aero.AeroScraper
	Edison  ed.EdisonScraper
	Ponrepo pon.PonrepoScraper
)

var AllCinemas = []m.CinemaScraper{
	Aero,
	Edison,
	Ponrepo,
}
