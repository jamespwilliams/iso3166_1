package iso3166

//go:generate go run github.com/jamespwilliams/iso3166/gen -o generated.go

type Country int64

func Alpha2(country Country) (alpha2 string) {
	return alpha2s[int(country)]
}
