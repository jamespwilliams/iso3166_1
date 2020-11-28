package iso3166

import (
	"math/rand"
	"testing"
	"time"

	"github.com/biter777/countries"
)

func BenchmarkOurAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		country := AllCountries[rand.Intn(len(AllCountries))]
		_ = Alpha2(country)
	}
}

var theirAllCountries = countries.All()

func BenchmarkTheirAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		country := theirAllCountries[rand.Intn(len(theirAllCountries))]
		_ = country.Alpha2()
	}
}

func BenchmarkOurAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		country := AllCountries[rand.Intn(len(AllCountries))]
		_ = Alpha3(country)
	}
}

func BenchmarkTheirAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		country := theirAllCountries[rand.Intn(len(theirAllCountries))]
		_ = country.Alpha3()
	}
}
