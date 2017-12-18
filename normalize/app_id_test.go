package normalize

import (
	"regexp"
	"strings"
	"testing"
)

func TestNormalizeAppId(t *testing.T) {
	iOSId := "628677149"
	if normalizedId := AppId(iOSId); normalizedId != iOSId {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", iOSId, normalizedId)
	}

	appsFlyerId := "id628677149"
	if normalizedId := AppId(appsFlyerId); normalizedId != iOSId {
		t.Errorf("Expected id prefix to be stripped from '%s' but it returned '%s'", appsFlyerId, normalizedId)
	}

	androidId := "id.indonesian.app"
	if normalizedId := AppId(androidId); normalizedId != androidId {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", androidId, normalizedId)
	}
}

var iosAppstoreRefRegexp = regexp.MustCompile(`^(?:id)[0-9]+$`)
var benchAppId = "id123456789f"

// AppId returns a normalized version of the app id.
// At the moment, it only removes the "id" prefix from iOS App Ids
// that AppsFyler adds.
func oldAppId(ref string) string {
	if iosAppstoreRefRegexp.MatchString(ref) {
		return strings.TrimLeft(ref, "id")
	}
	return ref
}

func BenchmarkAppIdRegex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			oldAppId(benchAppId)
		}
	})
}

func BenchmarkAppIdRagel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			AppId(benchAppId)
		}
	})
}
