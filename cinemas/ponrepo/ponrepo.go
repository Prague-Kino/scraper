package ponrepo

import "github.com/Prague-Kino/cast"

var Ponrepo = cast.Kino{
	Name:       "Ponrepo",
	BaseDomain: "nfa.cz",
	ProgramURL: "https://nfa.cz/en/ponrepo-cinema/program/program",
}

// Ponrepo doesn't have the pricing information
// on the program website, so it can't be scraped.
// But most of them cost 130 CZK.
const DefaultPrice = 130
