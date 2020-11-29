package iso3166_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromAlpha2(t *testing.T) {
	for _, c := range AllCountries {
		c2, err := FromAlpha2(c.Alpha2)
		assert.NoError(t, err)
		assert.Equal(t, c.Numeric, c2.Numeric)
	}
}

func TestFromAlpha3(t *testing.T) {
	for _, c := range AllCountries {
		c2, err := FromAlpha3(c.Alpha3)
		assert.NoError(t, err)
		assert.Equal(t, c.Numeric, c2.Numeric)
	}
}

func TestDirectAccessSmoke(t *testing.T) {
	assert.Equal(t, 826, UnitedKingdomOfGreatBritainAndNorthernIreland.Numeric)
}

func TestFromAlpha2Smoke(t *testing.T) {
	c2, err := FromAlpha2("GB")
	assert.NoError(t, err)
	assert.Equal(t, 826, c2.Numeric)
}

func TestFromAlpha3Smoke(t *testing.T) {
	c2, err := FromAlpha3("GBR")
	assert.NoError(t, err)
	assert.Equal(t, 826, c2.Numeric)
}
