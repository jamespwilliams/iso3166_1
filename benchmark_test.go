package iso3166_1

import (
	"math/rand"
	"testing"
	"time"
)

var (
	a2s []string
	a3s []string
)

func init() {
	for _, c := range AllCountries {
		a2s = append(a2s, c.Alpha2)
	}

	for _, c := range AllCountries {
		a3s = append(a3s, c.Alpha3)
	}
}

func BenchmarkLookupAlpha2Slice(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a2 := a2s[rand.Intn(len(a2s))]
		_, _ = FromAlpha2Slice(a2)
	}
}

func BenchmarkLookupAlpha3Slice(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		a3 := a3s[rand.Intn(len(a3s))]
		_, _ = FromAlpha2Slice(a3)
	}
}
