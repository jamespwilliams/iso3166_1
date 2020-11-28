// Package iso3166_1 provides efficient functions for accessing information about countries, as standardised in
// ISO-3166-1
package iso3166_1

//go:generate go run github.com/jamespwilliams/iso3166_1/gen -o generated.go

type Country int

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
