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
