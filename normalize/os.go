//go:generate ragel -Z -G2 os_fsm.rl
package normalize

import "strings"

// Os returns a normalized OS name (android, ios, tvos, roku, fireos, smartcast, webos, tizen, linux, vidaa, reachtv, or chromeos)
// based on what can be found in the given string
func Os(os string) string {
	if os == "" {
		return ""
	}
	// Check more specific patterns first to avoid false matches
	if MatchOsFireOS(os) {
		return "fireos"
	}
	if MatchOsSmartCast(os) {
		return "smartcast"
	}
	if MatchOsChromeOS(os) {
		return "chromeos"
	}
	if MatchOsTvOS(os) {
		return "tvos"
	}
	if MatchOsRoku(os) {
		return "roku"
	}
	if MatchOsWebOS(os) {
		return "webos"
	}
	if MatchOsTizen(os) {
		return "tizen"
	}
	if MatchOsVidaa(os) {
		return "vidaa"
	}
	if MatchOsReachTV(os) {
		return "reachtv"
	}
	if MatchOsiOS(os) {
		return "ios"
	}
	if MatchOsAndroid(os) {
		return "android"
	}
	if MatchOsLinux(os) {
		return "linux"
	}

	// For non-normalized OS values, remove spaces and convert to lowercase
	// to consolidate variations (e.g., "Unknown OS" and "UnknownOS" -> "unknownos")
	return strings.ToLower(strings.ReplaceAll(os, " ", ""))
}
