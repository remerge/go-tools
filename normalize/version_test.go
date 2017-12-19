package normalize

import (
	"testing"
)

func TestVersion(t *testing.T) {
	// Basic Test. Logic is tested in version.Version#String
	normalizedVersion := Version([]byte("android : 5.1.0 : L : sdk=21 : LOLLIPOP : sdk=21"))
	if normalizedVersion != "5.1" {
		t.Errorf("Expected normalized version to '5.1' but it was '%s'", normalizedVersion)
	}
}
