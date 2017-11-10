package normalize

import (
	"testing"
)

func TestCountryNormalizedAlpha2(t *testing.T) {
	for _, v := range [][2]string{
		{"de", "DE"},
		{"gb", "gbr"},
		{"us", "United States"},
		{"xx", "Arstotzka"},
		{"xx", ""},
	} {
		actual := CountryAlpha2(v[1])
		if actual != v[0] {
			t.Errorf("expected %s for '%s' but got %s", v[0], v[1], actual)
		}
	}
}
