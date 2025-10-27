//go:generate ragel -Z -G2 os_fsm.rl
package normalize

// Os returns either android, ios, or ctv depending
// on what can be found in the given string
func Os(os string) string {
	if os == "" {
		return ""
	}
	if MatchOsiOS(os) {
		return "ios"
	}
	if MatchOsCTV(os) {
		return "ctv"
	}
	if MatchOsAndroid(os) {
		return "android"
	}
	return ""
}
