package normalize

import (
	"regexp"
	"strings"
	"testing"
)

func TestNormalizeAppId(t *testing.T) {
	iOSId := "628677149"
	if normalizedID := AppId(iOSId); normalizedID != iOSId {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", iOSId, normalizedID)
	}

	appsFlyerID := "id628677149"
	if normalizedID := AppId(appsFlyerID); normalizedID != iOSId {
		t.Errorf("Expected id prefix to be stripped from '%s' but it returned '%s'", appsFlyerID, normalizedID)
	}

	androidID := "id.indonesian.app"
	if normalizedID := AppId(androidID); normalizedID != androidID {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", androidID, normalizedID)
	}
}

var iosAppstoreRefRegexp = regexp.MustCompile(`^(?:id)[0-9]+$`)
var benchAppID = "id123456789f"

// AppId returns a normalized version of the app id.
// At the moment, it only removes the "id" prefix from iOS App Ids
// that AppsFyler adds.
func oldAppID(ref string) string {
	if iosAppstoreRefRegexp.MatchString(ref) {
		return strings.TrimLeft(ref, "id")
	}
	return ref
}

func BenchmarkAppIdRegex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			oldAppID(benchAppID)
		}
	})
}

func BenchmarkAppIdRagel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = AppId(benchAppID)
		}
	})
}
