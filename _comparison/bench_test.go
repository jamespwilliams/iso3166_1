package main

import (
	"math/rand"
	"testing"
	"time"

	bookpeople "github.com/TheBookPeople/iso3166"
	biter777 "github.com/biter777/countries"
	"github.com/jamespwilliams/iso3166_1"
	launchdarkly "github.com/launchdarkly/go-country-codes"
	"github.com/pariz/gountries"
)

var (
	a2s            []string
	a3s            []string
	gountriesQuery *gountries.Query
)

func init() {
	for _, c := range iso3166_1.AllCountries {
		a2s = append(a2s, c.Alpha2)
	}

	for _, c := range iso3166_1.AllCountries {
		a3s = append(a3s, c.Alpha3)
	}

	gountriesQuery = gountries.New()
}

func BenchmarkOurFromAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_, _ = iso3166_1.FromAlpha2(a2)
	}
}

func BenchmarkOurFromAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_, _ = iso3166_1.FromAlpha3(a3)
	}
}

func BenchmarkBiter777FromAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_ = biter777.ByName(a2)
	}
}

func BenchmarkBiter777FromAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_ = biter777.ByName(a3)
	}
}

func BenchmarkBookPeopleFromAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_, _ = bookpeople.Decode(a2, "", false)
	}
}

func BenchmarkBookPeopleFromAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_, _ = bookpeople.Decode(a3, "", false)
	}
}

func BenchmarkLaunchDarklyFromAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_, _ = launchdarkly.GetByAlpha2(a2)
	}
}

func BenchmarkLaunchDarklyFromAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_, _ = launchdarkly.GetByAlpha3(a3)
	}
}

func BenchmarkGountriesFromAlpha2(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_, _ = gountriesQuery.FindCountryByAlpha(a2)
	}
}

func BenchmarkGountriesFromAlpha3(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_, _ = gountriesQuery.FindCountryByAlpha(a3)
	}
}
