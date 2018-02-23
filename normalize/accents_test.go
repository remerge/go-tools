package normalize

import (
	"testing"
)

func TestRemoveAccents(t *testing.T) {
	for _, v := range [][2]string{
		{"Smøre Brø", "Smoere Broe"},
		{"Müh Pö Pä", "Muh Po Pa"},
		{"ÇÆśß", "CAEsss"},
	} {
		actual := RemoveAccents(v[0])
		if actual != v[1] {
			t.Errorf("expected %s for '%s' but got %s", v[1], v[0], actual)
		}
	}
}
