//go:generate ragel -Z -G2 uuid_fsm.rl

package uuid

import (
	"github.com/remerge/go-tools/normalize"
)

// isValidFast checks the format of the provided UUID string for valid length and
// delimiting dashes. Only a structural check!
func isValidFast(uuid []byte) bool {
	return len(uuid) == 36 && uuid[8] == '-' && uuid[13] == '-' && uuid[18] == '-'
}

// IsValid return true if uuid matches the UUID standard regexp
func IsValid(uuid []byte) bool {
	return isValidFast(uuid) && matchUuidRegex(uuid)
}

// IsiOS returns true if uuid matches the iOS specifc UUID regexp
func IsiOS(uuid []byte) bool {
	return isValidFast(uuid) && matchUuidRegexiOS(uuid)
}

// IsAndroid returns true if uuid matches the Android specifc UUID regexp
func IsAndroid(uuid []byte) bool {
	return isValidFast(uuid) && matchUuidRegexAndroid(uuid)
}

func MatchesOS(expectedOS string, givenOS string, id []byte) bool {
	if expectedOS == "" {
		return true
	}
	if givenOS != "" && normalize.Os(givenOS) != expectedOS {
		return false
	}
	switch expectedOS {
	case "ios":
		return IsiOS(id)
	case "android":
		return IsAndroid(id)
	}
	return true
}
