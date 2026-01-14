// There are four cinemas that are all part of the Aerokina group
// We can scrape them all from the same URL and differentiate by cinema name

package aero

import (
	"github.com/Prague-Kino/scraper/lib"
)

const (
	BASE_DOMAIN      = "www.kinoaero.cz"
	BASE_PROGRAM_URL = "https://www.kinoaero.cz/en?cinema=1%2C2%2C9%2C3&sort=sort-by-data&hall=34%2C35%2C1%2C2%2C3%2C24&english-friendly=1"
)

var (
	Aero     = newAerokindaCinema("Aero")
	Svetozor = newAerokindaCinema("Světozor")
	Lucerna  = newAerokindaCinema("Lucerna")
	BigOko   = newAerokindaCinema("Big Oko")
)

func newAerokindaCinema(name string) *lib.Kino {
	k := lib.Kino{
		Name:       name,
		BaseDomain: BASE_DOMAIN,
		ProgramURL: BASE_PROGRAM_URL,
	}
	return &k
}
