package normalize

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	for _, test := range [][3]string{
		{"$", "-", "USD"},
		{"€", "-", "EUR"},
		{"eur", "-", "EUR"},
		{"uSd", "-", "USD"},
		{"rub", "-", "RUB"},
		{"narf", "-", "-"},
		{"", "EUR", "EUR"},
		{"blab", "EUR", "EUR"},
		{"asdasd", "", ""},
	} {
		r := Currency(test[0], test[1])
		if r != test[2] {
			t.Errorf("The normalized currency for '%s' should be '%s' but is '%s'", test[0], test[2], r)
		}
	}
}
