package normalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	langs := [][2]string{
		{"en", "en"},                            // ok
		{"EN", "en"},                            // upper case
		{"de-de", "de"},                         // dash
		{"EN_us", "en"},                         // another upper case
		{"ko_kr", "ko"},                         // underscore
		{"i", ""},                               // ???
		{"I", ""},                               // ???
		{"zh-hant-tw", "zh"},                    // ???
		{"ko-kr,ja-jp;q=0.9,en-us;q=0.8", "ko"}, // bad
		{"undefined", ""},                       // bad
		{"-s", ""},                              // ???
		{"0", ""},                               // ???
		// terminals
		{"af-af","af"},
		{"AF-AF", "af"},
		{"zu-zu", "zu"},
		{"ZU-ZU", "zu"},
		{"bad", ""},
	}

	for _, l := range langs {
		assert.Equal(t, l[1], Language(l[0]))
	}
}
