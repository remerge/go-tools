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
		"tvos": {
			"tvOS",
			"TVOS",
			"tvos",
			"Apple tvOS",
		},
		"roku": {
			"roku",
			"Roku",
			"ROKU",
			"rokuos",
			"RokuOS",
			"roku os",
			"roku+os",
			"Roku+OS",
		},
		"fireos": {
			"fireos",
			"FireOS",
			"FIREOS",
			"fire-os",
			"fire+os",
			"Fire+OS",
			"fire",
			"Fire",
			"FIRE",
			"firetvos",
			"FireTVOS",
		},
		"smartcast": {
			"smartcast",
			"SmartCast",
			"SMARTCAST",
			"smartcastos",
			"SmartCastOS",
			"smartcast%20os",
			"SmartCast%20OS",
			"viziosmartcast",
			"VizioSmartCast",
			"VIZIOSMARTCAST",
		},
		"webos": {
			"webos",
			"WebOS",
			"WEBOS",
			"webostv",
			"WebOSTV",
			"webos tv",
			"webos+tv",
			"WebOS+TV",
		},
		"tizen": {
			"tizen",
			"Tizen",
			"TIZEN",
			"tizentv",
			"TizenTV",
			"TIZENTV",
		},
		"linux": {
			"linux",
			"Linux",
			"LINUX",
		},
		"vidaa": {
			"vidaa",
			"Vidaa",
			"VIDAA",
		},
		"reachtv": {
			"reachtv",
			"ReachTV",
			"REACHTV",
		},
		"chromeos": {
			"chromeos",
			"ChromeOS",
			"CHROMEOS",
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
