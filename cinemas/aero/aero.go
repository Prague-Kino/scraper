package aero

import (
	"prague-kino/lib"
)

const (
	KINO_NAME   = "Aero"
	BASE_DOMAIN = "www.kinoaero.cz"
	PROGRAM_URL = "https://www.kinoaero.cz/en?cinema=1&sort=sort-by-data&english-friendly=1"
)

var Aero = lib.Kino{
	Name:       KINO_NAME,
	BaseDomain: "www.kinoaero.cz",
	ProgramURL: PROGRAM_URL,
}
