package normalize

import (
	"testing"
)

func TestNormalizeOs(t *testing.T) {
	for os, strings := range map[string][]string{
		"android": {
			"Android OS",
			"  some anDroid OS",
			"android",
			"ANDROID",
			"	anDroid_v1",
		},
		"ios": {
			"iPhone X",
			"   IPAD  ",
			" MiniIpad",
			"iOS",
			"			ios",
			"iPod with ios",
		},
	} {
		for _, s := range strings {
			nos := Os(s)
			if nos != os {
				t.Errorf("expected %s for '%s' but got %s", os, s, nos)
			}
		}
	}
}
