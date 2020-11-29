// Package iso3166_1 provides efficient functions for accessing information about countries, as standardised in
// ISO-3166-1
package iso3166_1

import "errors"

//go:generate go run github.com/jamespwilliams/iso3166_1/gen -o generated.go

type Country int

var ErrNoSuchCountry = errors.New("iso3166_1: no such country exists")

func FromNumeric(c int) Country {
	return Country(c)
}

func FromAlpha2(a2 string) (Country, error) {
	if c, ok := alpha2Lookup[a2]; ok {
		return c, nil
	}

	return 0, ErrNoSuchCountry
}

func FromAlpha2Slice(a2 string) (Country, error) {
	f := uint16(a2[0] - 0x41)
	s := uint16(a2[1] - 0x41)

	if f > 25 {
		f -= 0x20
		if f > 25 {
			return 0, ErrNoSuchCountry
		}
	}

	if s > 25 {
		s -= 0x20
		if s > 25 {
			return 0, ErrNoSuchCountry
		}
	}

	index := f<<5 + s

	c := alpha2SliceLookup[index]
	if c == 0 {
		return 0, ErrNoSuchCountry
	}

	return Country(c), nil
}

func FromAlpha3Slice(a3 string) (Country, error) {
	f := uint16(a3[0] - 0x41)
	s := uint16(a3[1] - 0x41)
	t := uint16(a3[2] - 0x41)

	if f > 25 {
		f -= 0x20
		if f > 25 {
			return 0, ErrNoSuchCountry
		}
	}

	if s > 25 {
		s -= 0x20
		if s > 25 {
			return 0, ErrNoSuchCountry
		}
	}

	if t > 25 {
		t -= 0x20
		if t > 25 {
			return 0, ErrNoSuchCountry
		}
	}

	index := f<<10 + s<<5 + t
	c := alpha3SliceLookup[index]
	if c == 0 {
		return 0, ErrNoSuchCountry
	}

	return Country(c), nil
}

// Numeric returns the Country's Numeric code, as defined in ISO3166-1
func (c Country) Numeric() int {
	return int(c)
}

// Alpha2 returns the Country's Alpha-2 code, as defined in ISO3166-1
func (c Country) Alpha2() string {
	return alpha2s[c]
}

// Alpha3 returns the Country's Alpha-3 code, as defined in ISO3166-1
func (c Country) Alpha3() string {
	return alpha3s[c]
}

// Name returns the Country's official "short name", as defined in ISO3166-1
func (c Country) Name() string {
	return shortNames[c]
}

func (c Country) String() string {
	return c.Name()
}
