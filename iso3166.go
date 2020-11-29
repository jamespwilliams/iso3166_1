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
	firstRune := a2[0]
	secondRune := a2[1]

	firstRune -= 0x41
	secondRune -= 0x41

	fr := uint16(firstRune)
	sr := uint16(secondRune)

	index := fr<<5 + sr

	c := alpha2SliceLookup[index]
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
