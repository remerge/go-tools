//go:generate ragel -Z -G2 uuid_fsm.rl

package uuid

import (
	"github.com/remerge/go-tools/normalize"
)

// isValidFast checks the format of the provided UUID string for valid length and
// delimiting dashes. Only a structural check!
func isValidFast(uuid string) bool {
	return len(uuid) == 36 && uuid[8] == '-' && uuid[13] == '-' && uuid[18] == '-'
}

// IsValid return true if uuid matches the UUID standard regexp
func IsValid(uuid string) bool {
	return matchUuidRegex([]byte(uuid))
}

// IsiOS returns true if uuid matches the iOS specifc UUID regexp
func IsiOS(uuid string) bool {
	if !isValidFast(uuid) {
		return false
	}
	return matchUuidRegexiOS([]byte(uuid))
}

// IsAndroid returns true if uuid matches the Android specifc UUID regexp
func IsAndroid(uuid string) bool {
	if !isValidFast(uuid) {
		return false
	}
	return matchUuidRegexAndroid([]byte(uuid))
}

func MatchesOS(expectedOS string, givenOS string, id string) bool {
	if expectedOS == "" {
		return true
	}
	if givenOS != "" && normalize.Os(givenOS) != expectedOS {
		return false
	}
	switch expectedOS {
	case "ios":
		if !IsiOS(id) {
			return false
		}
	case "android":
		if !IsAndroid(id) {
			return false
		}
	}
	return true
}
