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
		"ctv": {
			"tvOS",
			"tvOS 15.0",
			"Apple tvOS",
			"tvOS_16",
			"tv_os",
			"Android TV",
			"ANDROID TV",
			"android_tv",
			"Android_TV_12",
			"webOS",
			"webOS 3.0",
			"web_os",
			"webOS_5.0",
			"Tizen",
			"TIZEN",
			"tizen_v6",
			"Roku",
			"roku_os",
			"Fire TV",
			"FIRE TV",
			"fire_tv",
			"Fire_TV_Stick",
			"vidaa",
			"VIDAA",
			"vidaa_u5",
			"vewd",
			"smartcast",
			"SmartCast_TV",
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
var benchOs = "iPhone X"

func BenchmarkOsRegex(b *testing.B) {
	var hits int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if osRegExp.MatchString(benchOs) {
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
