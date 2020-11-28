package iso3166

//go:generate go run github.com/jamespwilliams/iso3166/gen -o generated.go

type Country int64

func (c Country) Alpha2() (alpha2 string) {
	return alpha2s[int(c)]
}

func (c Country) Alpha3() (alpha3 string) {
	return alpha3s[int(c)]
}
