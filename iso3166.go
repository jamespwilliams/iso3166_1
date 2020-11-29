// Package iso3166_1 provides efficient functions for accessing information about countries, as standardised in
// ISO-3166-1
package iso3166_1

import "errors"

//go:generate go run github.com/jamespwilliams/iso3166_1/gen -o generated.go

type Country struct {
	Numeric int
	Alpha2  string
	Alpha3  string
	Name    string
}

var ErrNoSuchCountry = errors.New("iso3166_1: no such country exists")

func FromNumeric(c int) Country {
	return structSlice[c]
}

func FromAlpha2(a2 string) (Country, error) {
	f := uint16(a2[0] - 0x41)
	s := uint16(a2[1] - 0x41)

	if f > 25 {
		f -= 0x20
		if f > 25 {
			return Country{}, ErrNoSuchCountry
		}
	}

	if s > 25 {
		s -= 0x20
		if s > 25 {
			return Country{}, ErrNoSuchCountry
		}
	}

	index := f<<5 + s

	c := alpha2SliceLookup[index]
	if c == 0 {
		return Country{}, ErrNoSuchCountry
	}

	return structSlice[c], nil
}

func FromAlpha3(a3 string) (Country, error) {
	f := uint16(a3[0] - 0x41)
	s := uint16(a3[1] - 0x41)
	t := uint16(a3[2] - 0x41)

	if f > 25 {
		f -= 0x20
		if f > 25 {
			return Country{}, ErrNoSuchCountry
		}
	}

	if s > 25 {
		s -= 0x20
		if s > 25 {
			return Country{}, ErrNoSuchCountry
		}
	}

	if t > 25 {
		t -= 0x20
		if t > 25 {
			return Country{}, ErrNoSuchCountry
		}
	}

	index := f<<10 + s<<5 + t
	c := alpha3SliceLookup[index]
	if c == 0 {
		return Country{}, ErrNoSuchCountry
	}

	return structSlice[c], nil
}
