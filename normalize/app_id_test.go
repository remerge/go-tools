package normalize

import (
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
