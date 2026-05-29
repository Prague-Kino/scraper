// There are four cinemas that are all part of the Aerokina group
// We can scrape them all from the same URL and differentiate by cinema name

package aero

import (
	"strings"

	"github.com/Prague-Kino/cast"
)

const (
	BASE_DOMAIN = "www.kinoaero.cz"
	PROGRAM_URL = "https://www.kinoaero.cz/en?cinema=1%2C2%2C9%2C7%2C3&sort=sort-by-data&hall=34%2C35%2C1%2C2%2C3%2C24&english-friendly=1"
)

var (
	Aero       = newAerokinaCinema("Kino Aero")
	Svetozor   = newAerokinaCinema("Kino Světozor")
	Lucerna    = newAerokinaCinema("Kino Lucerna")
	BigOko     = newAerokinaCinema("Kino Big Oko")
	Pritomnost = newAerokinaCinema("Kino Pritomnost")
)

var aerokinaMap = map[string]*cast.Kino{
	"Aero":       Aero,
	"Světozor":   Svetozor,
	"Lucerna":    Lucerna,
	"Bio Oko":    BigOko,
	"Přítomnost": Pritomnost,
}

func newAerokinaCinema(name string) *cast.Kino {
	return cast.NewKino(name, BASE_DOMAIN, PROGRAM_URL)
}

func getKinoByName(name string) cast.Kino {
	kino, ok := aerokinaMap[strings.TrimSpace(name)]
	if ok {
		return *kino
	}

	return *Aero
}
