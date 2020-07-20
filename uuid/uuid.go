//go:generate ragel -Z -G2 uuid_fsm.rl

package uuid

import (
	"github.com/remerge/go-tools/normalize"
)

// IsValidFast checks the format of the provided UUID string for valid length
// and delimiting dashes. Only a structural check!
func IsValidFast(uuid string) bool {
	return len(uuid) == 36 && uuid[8] == '-' && uuid[13] == '-' && uuid[18] == '-'
}

// IsValid return true if uuid matches the UUID standard regexp
func IsValid(uuid string) bool {
	return IsValidFast(uuid) && matchUuidRegex(uuid) && !isTest(uuid)
}

// IsiOS returns true if uuid matches the iOS specifc UUID regexp
func IsiOS(uuid string) bool {
	return IsValidFast(uuid) && matchUuidRegexiOS(uuid) && !isTest(uuid)
}

// IsAndroid returns true if uuid matches the Android specifc UUID regexp
func IsAndroid(uuid string) bool {
	return IsValidFast(uuid) && matchUuidRegexAndroid(uuid) && !isTest(uuid)
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
		return IsiOS(id)
	case "android":
		return IsAndroid(id)
	}
	return true
}

func isTest(uuid string) bool {
	return uuid == "00000000-0000-0000-0000-000000000000" || uuid == "ffffffff-ffff-ffff-ffff-ffffffffffff" || uuid == "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF"
}
