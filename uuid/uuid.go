package uuid

import (
	"regexp"

	"github.com/remerge/go-tools/normalize"
)

var uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
var uuidRegexiOS = regexp.MustCompile("^[A-F0-9]{8}-[A-F0-9]{4}-4[A-F0-9]{3}-[8|9|A|B][A-F0-9]{3}-[A-F0-9]{12}$")
var uuidRegexAndroid = regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|a|b][a-f0-9]{3}-[a-f0-9]{12}$")

// IsValidFast checks the format of the provided UUID string for valid length and
// delimiting dashes. Only a structural check!
func IsValidFast(uuid string) bool {
	return len(uuid) == 36 && uuid[8] == '-' && uuid[13] == '-' && uuid[18] == '-'
}

// IsValid return true iff uuid matches the UUID standard regexp
func IsValid(uuid string) bool {
	return uuidRegex.MatchString(uuid)
}

// IsiOS returns true if uuid matches the iOS specifc UUID regexp
func IsiOS(uuid string) bool {
	if !IsValidFast(uuid) {
		return false
	}
	return uuidRegexiOS.MatchString(uuid)
}

// IsAndroid returns true if uuid matches the Android specifc UUID regexp
func IsAndroid(uuid string) bool {
	if !IsValidFast(uuid) {
		return false
	}
	return uuidRegexAndroid.MatchString(uuid)
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
