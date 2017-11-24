package normalize

import (
	"testing"
)

func TestNormalizeAppID(t *testing.T) {
	iOSID := "628677149"
	if noramlizedID := AppID(iOSID); noramlizedID != iOSID {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", iOSID, noramlizedID)
	}

	appsFlyerID := "id628677149"
	if noramlizedID := AppID(appsFlyerID); noramlizedID != iOSID {
		t.Errorf("Expected id prefix to be stripped from '%s' but it returned '%s'", appsFlyerID, noramlizedID)
	}

	androidID := "id.indonesian.app"
	if noramlizedID := AppID(androidID); noramlizedID != androidID {
		t.Errorf("Expected ID '%s' to remain unchanged but it was '%s'", androidID, noramlizedID)
	}
}
