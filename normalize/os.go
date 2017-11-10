package normalize

import "regexp"

var osRegExp = regexp.MustCompile("(?i)(?P<ios>ipod|iphone|ipad|ios)|(?P<android>android)")

// Os returns either android or ios depending
// on what can be found in the given string
func Os(os string) string {
	if os == "" {
		return ""
	}
	match := osRegExp.FindStringSubmatch(os)
	if len(match) == 0 {
		return ""
	}
	if len(match[1]) > 0 {
		return "ios"
	}
	if len(match[2]) > 0 {
		return "android"
	}
	return ""
}
