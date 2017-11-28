package normalize

import (
	"regexp"
	"strings"
)

var iosAppstoreRefRegexp = regexp.MustCompile(`^(?:id)[0-9]+$`)

// AppId returns a normalized version of the app id.
func AppId(ref string) string {
	if iosAppstoreRefRegexp.MatchString(ref) {
		return strings.TrimLeft(ref, "id")
	}

	return ref
}
