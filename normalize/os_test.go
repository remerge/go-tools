package normalize

import (
	"regexp"
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

var osRegExp = regexp.MustCompile("(?i)(?P<ios>ipod|iphone|ipad|ios)|(?P<android>android)")
var benchOs = []byte("iPhone X")

func BenchmarkOsRegex(b *testing.B) {
	var hits int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if osRegExp.Match(benchOs) {
				hits++
			}
		}
	})
}

func BenchmarkOsRagel(b *testing.B) {
	var hits int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if MatchOsiOS(benchOs) {
				hits++
			}
		}
	})
}
