//go:generate ragel -Z -G2 os_fsm.rl
package normalize

// Os returns either android or ios depending
// on what can be found in the given string
func Os(os string) string {
	if os == "" {
		return ""
	}
	osBytes := []byte(os)
	if MatchOsiOS(osBytes) {
		return "ios"
	}
	if MatchOsAndroid(osBytes) {
		return "android"
	}
	return ""
}
