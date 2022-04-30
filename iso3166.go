// Package iso3166_1 provides efficient functions for accessing information about countries, as standardised in
// ISO-3166-1
package iso3166_1

import (
	"errors"
)

//go:generate go run github.com/jamespwilliams/iso3166_1/internal/gen -o generated.go

type Country struct {
	Numeric int
	Alpha2  string
	Alpha3  string
	Name    string
}

var ErrNoSuchCountry = errors.New("iso3166_1: no such country exists")

func FromNumeric(c int) *Country {
	return &countries[c]
}

func FromAlpha2(a2 string) (*Country, error) {
	f := uint16(a2[0] & 0b00011111)
	s := uint16(a2[1] & 0b00011111)

	index := f<<5 + s

	res := alpha2Lookup[index]
	if res == nil {
		return nil, ErrNoSuchCountry
	}

	return res, nil
}

func FromAlpha3(a3 string) (*Country, error) {
	f := uint16(a3[0] & 0b00011111)
	s := uint16(a3[1] & 0b00011111)
	t := uint16(a3[2] & 0b00011111)

	index := f<<10 + s<<5 + t

	res := alpha3Lookup[index]
	if res == nil {
		return nil, ErrNoSuchCountry
	}

	return res, nil
}
